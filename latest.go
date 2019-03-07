package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mholt/archiver"

	"github.com/taskie/osplus"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tcnksm/go-latest"
)

// https://golang.hateblo.jp/entry/2018/10/19/180000
type Progress struct {
	total int64
	size  int64
}

// io.Writer の要件を満たす
func (p *Progress) Write(data []byte) (int, error) {
	n := len(data)
	p.size += int64(n)
	fmt.Fprintf(os.Stderr, "\r%6d KiB / %6d KiB (%5.1f%%)", p.size/1024, p.total/1024, float64(p.size)/float64(p.total)*100)
	return n, nil
}

func downloadApplication(owner, repo, version, dir string) (fpath string, err error) {
	goOs := runtime.GOOS
	goArch := runtime.GOARCH
	archiveName := fmt.Sprintf("%s_%s_%s_%s.tar.gz", repo, version, goOs, goArch)
	// https://github.com/taskie/ore/releases/download/v0.1.2/ore_0.1.2_linux_amd64.tar.gz
	urlFormat := "https://github.com/%s/%s/releases/download/v%s/%s"
	url := fmt.Sprintf(urlFormat, owner, repo, version, archiveName)
	log.Infof("downloading: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	progress := Progress{}
	progress.total = resp.ContentLength
	if dir == "" {
		f, err := ioutil.TempFile("", repo+"-")
		if err != nil {
			return "", err
		}
		defer f.Close()
		fpath := f.Name()
		bw := bufio.NewWriter(f)
		defer bw.Flush()
		_, err = io.Copy(bw, io.TeeReader(resp.Body, &progress))
		if err != nil {
			return fpath, err
		}
		return fpath, nil
	}
	fpath = filepath.Join(dir, archiveName)
	log.Infof("destination: %s", fpath)
	w, commit, err := osplus.NewOpener().CreateTempFileWithDestination(fpath, "", repo+"-")
	if err != nil {
		return
	}
	defer w.Close()
	_, err = io.Copy(w, io.TeeReader(resp.Body, &progress))
	if err != nil {
		return
	}
	fmt.Fprintln(os.Stderr, "")
	commit(true)
	return
}

func extractApplication(commandName, fpath string) (err error) {
	log.Infof("extracting: %s", fpath)
	goPath, err := osplus.GetGoPath()
	if err != nil {
		return
	}
	destination := filepath.Join(goPath, "bin", commandName)
	w, commit, err := osplus.NewOpener().CreateTempFileWithDestination(destination, "", commandName+"-")
	if err != nil {
		return
	}
	defer w.Close()
	found := false
	var fileMode os.FileMode
	tgz := archiver.NewTarGz()
	err = tgz.Walk(fpath, func(f archiver.File) error {
		if f.Name() != commandName {
			return nil
		}
		fileMode = f.FileInfo.Mode()
		found = true
		_, err = io.Copy(w, f)
		return err
	})
	if err != nil {
		return
	}
	if !found {
		return fmt.Errorf("%s is not found in %s", commandName, fpath)
	}
	commit(true)
	w.Close()
	if fileMode == 0 {
		return fmt.Errorf("file mode must not be 0")
	}
	return os.Chmod(destination, fileMode)
}

func generateLatestCommand(owner, repo, currentVersion string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "latest",
	}
	githubTag := &latest.GithubTag{
		Owner:             owner,
		Repository:        repo,
		FixVersionStrFunc: latest.DeleteFrontV(),
	}
	cmd.AddCommand(&cobra.Command{
		Use: "status",
		Run: func(cmd *cobra.Command, args []string) {
			if currentVersion == "" {
				currentVersion = "0.0.0"
			}
			res, _ := latest.Check(githubTag, currentVersion)
			if res.Outdated {
				fmt.Printf("%s is not latest, you should upgrade to v%s", currentVersion, res.Current)
			} else {
				fmt.Printf("%s is latest now", currentVersion)
			}

		}})
	cmd.AddCommand(&cobra.Command{
		Use: "download",
		Run: func(cmd *cobra.Command, args []string) {
			if currentVersion == "" {
				log.Warn("version information is not embedded in this application.")
				return
			}
			if len(args) == 1 {
				fpath := args[0]
				err := extractApplication(repo, fpath)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			res, _ := latest.Check(githubTag, currentVersion)
			if res.Outdated {
				dir, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				_, err = downloadApplication(owner, repo, res.Current, dir)
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	})
	cmd.AddCommand(&cobra.Command{
		Use: "install",
		Run: func(cmd *cobra.Command, args []string) {
			if currentVersion == "" {
				log.Warn("version information is not embedded in this application.")
				return
			}
			if len(args) == 1 {
				fpath := args[0]
				err := extractApplication(repo, fpath)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			res, _ := latest.Check(githubTag, currentVersion)
			if res.Outdated {
				fpath, err := downloadApplication(owner, repo, res.Current, "")
				if err != nil {
					log.Fatal(err)
				}
				err = extractApplication(repo, fpath)
				if err != nil {
					log.Fatal(err)
				}
				_ = os.Remove(fpath)
			}
		}})
	return cmd
}
