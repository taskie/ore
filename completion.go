package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func generateCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "completion",
		Run: func(cmd *cobra.Command, args []string) {
			shell := ""
			if len(args) > 0 {
				shell = args[0]
			} else {
				shell = os.Getenv("SHELL")
				if shell == "" {
					shell = "bash"
				}
			}
			ss := strings.SplitN(shell, " ", 2)
			base := filepath.Base(ss[0])
			switch strings.ToLower(base) {
			case "zsh":
				Command.GenZshCompletion(os.Stdout)
			default:
				Command.GenBashCompletion(os.Stdout)
			}
		},
	}
	return cmd
}
