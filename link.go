package main

import (
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/taskie/osplus"
)

func generateLinkCommand(subcommandNames []string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "link",
		Run: func(cmd *cobra.Command, args []string) {
			gopath, err := osplus.GetGoPath()
			if err != nil {
				log.Fatal(err)
			}
			for _, name := range subcommandNames {
				dst := filepath.Join(gopath, "bin", name)
				abs, err := exec.LookPath(os.Args[0])
				if err != nil {
					log.Warn(err)
					continue
				}
				ln, err := os.Readlink(dst)
				if err == nil {
					// link exists
					if ln == abs {
						continue
					} else if force, err := cmd.Flags().GetBool("force"); err == nil && force {
						os.Remove(dst)
					} else {
						log.Warnf("%s is not %s, but %s", dst, abs, ln)
						continue
					}
				}
				err = os.Symlink(abs, dst)
				if err != nil {
					log.Warn(err)
					continue
				}
			}
		},
	}
	cmd.Flags().BoolP("force", "f", false, "link force")
	return cmd
}

func generateUnlinkCommand(subcommandNames []string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "unlink",
		Run: func(cmd *cobra.Command, args []string) {
			gopath, err := osplus.GetGoPath()
			if err != nil {
				log.Fatal(err)
			}
			for _, name := range subcommandNames {
				dst := filepath.Join(gopath, "bin", name)
				abs, err := exec.LookPath(os.Args[0])
				if err != nil {
					log.Warn(err)
					continue
				}
				ln, err := os.Readlink(dst)
				if err != nil {
					log.Warn(err)
					continue
				}
				if ln != abs {
					if force, err := cmd.Flags().GetBool("force"); err == nil && !force {
						log.Warnf("%s is not %s, but %s", dst, abs, ln)
						continue
					}
				}
				err = os.Remove(dst)
				if err != nil {
					log.Warn(err)
					continue
				}
			}
		},
	}
	cmd.Flags().BoolP("force", "f", false, "unlink force")
	return cmd
}
