package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/taskie/csvt/cli/csvt"
	"github.com/taskie/fwv/cli/fwv"
	"github.com/taskie/gfp/cli/gfp"
	"github.com/taskie/gtp/cli/gtp"
	"github.com/taskie/jc/cli/jc"
	"github.com/taskie/reinc/cli/reinc"
	"github.com/taskie/rlexec/cli/rlexec"
)

var (
	version, commit, date string
	verbose, debug        bool
)

const (
	Owner       = "taskie"
	CommandName = "ore"
)

var Command = &cobra.Command{
	Use: CommandName,
}

var SubcommandNames []string

func init() {
	Command.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	Command.PersistentFlags().BoolVar(&verbose, "debug", false, "debug output")

	Command.AddCommand(jc.Command)
	Command.AddCommand(gtp.Command)
	Command.AddCommand(csvt.Command)
	Command.AddCommand(fwv.Command)
	Command.AddCommand(reinc.Command)
	Command.AddCommand(rlexec.Command)
	Command.AddCommand(gfp.Command)
	addExtraCommands()
	SubcommandNames := make([]string, 0)
	for _, cmd := range Command.Commands() {
		SubcommandNames = append(SubcommandNames, cmd.Name())
	}
	Command.AddCommand(generateVersionCommand())
	Command.AddCommand(&cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			for _, name := range SubcommandNames {
				fmt.Println(name)
			}
		},
	})
	Command.AddCommand(generateLinkCommand(SubcommandNames))
	Command.AddCommand(generateUnlinkCommand(SubcommandNames))
	Command.AddCommand(generateCompletionCommand())
	Command.AddCommand(generateLatestCommand(Owner, CommandName, version))
}

func main() {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
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
