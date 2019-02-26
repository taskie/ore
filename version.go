package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func generateVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				switch strings.ToLower(args[0]) {
				case "commit", "c":
					fmt.Println(commit)
				case "date", "d":
					fmt.Println(date)
				default:
					fmt.Println(version)
				}
				return
			}
			if verbose, err := cmd.Flags().GetBool("verbose"); err == nil && verbose {
				fmt.Println("Version : " + version)
				fmt.Println("Commit  : " + commit)
				fmt.Println("Date    : " + date)
				return
			}
			fmt.Println(version)
		},
	}
	cmd.Flags().BoolP("verbose", "v", false, "verbose output")
	return cmd
}
