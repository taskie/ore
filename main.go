package main

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/taskie/csvt/cli/csvt"
	"github.com/taskie/fwv/cli/fwv"
	"github.com/taskie/gfp/cli/gfp"
	"github.com/taskie/gtp/cli/gtp"
	"github.com/taskie/jc/cli/jc"
	"github.com/taskie/osplus"
	"github.com/taskie/pity/cli/pity"
	"github.com/taskie/reinc/cli/reinc"
	"github.com/taskie/rltee/cli/rltee"
)

var (
	version  string
	revision string
)

const (
	CommandName = "ore"
)

var Command = &cobra.Command{
	Use: CommandName,
}

func init() {
	Command.AddCommand(jc.Command)
	Command.AddCommand(gtp.Command)
	Command.AddCommand(csvt.Command)
	Command.AddCommand(fwv.Command)
	Command.AddCommand(reinc.Command)
	Command.AddCommand(pity.Command)
	Command.AddCommand(rltee.Command)
	Command.AddCommand(gfp.Command)

	cmdNames := make([]string, 0)
	for _, cmd := range Command.Commands() {
		cmdNames = append(cmdNames, cmd.Name())
	}

	Command.AddCommand(&cobra.Command{
		Use: "link",
		Run: func(cmd *cobra.Command, args []string) {
			gopath, err := osplus.GetGoPath()
			if err != nil {
				log.Fatal(err)
			}
			for _, cmdName := range cmdNames {
				dst := filepath.Join(gopath, "bin", cmdName)
				abs, err := filepath.Abs(os.Args[0])
				if err != nil {
					log.Warn(err)
					continue
				}
				err = os.Symlink(abs, dst)
				if err != nil {
					log.Warn(err)
				}
			}
		},
	})
	Command.AddCommand(&cobra.Command{
		Use: "unlink",
		Run: func(cmd *cobra.Command, args []string) {
			gopath, err := osplus.GetGoPath()
			if err != nil {
				log.Fatal(err)
			}
			for _, cmdName := range cmdNames {
				dst := filepath.Join(gopath, "bin", cmdName)
				err = os.Remove(dst)
				if err != nil {
					log.Warn(err)
				}
			}
		},
	})
	Command.AddCommand(&cobra.Command{
		Use: "bash",
		Run: func(cmd *cobra.Command, args []string) {
			Command.GenBashCompletion(os.Stdout)
		},
	})
	Command.AddCommand(&cobra.Command{
		Use: "zsh",
		Run: func(cmd *cobra.Command, args []string) {
			Command.GenZshCompletion(os.Stdout)
		},
	})
}

func main() {
	base := filepath.Base(os.Args[0])
	if base == CommandName {
		err := Command.Execute()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		cmd, ss, err := Command.Find(append([]string{base}, os.Args[1:]...))
		if err != nil {
			log.Fatal(err)
		}
		os.Args = append([]string{CommandName, base}, ss...)
		err = cmd.Execute()
		if err != nil {
			log.Fatal(err)
		}
	}
}
