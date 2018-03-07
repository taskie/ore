package main

import (
	"github.com/taskie/gtp"
	"github.com/taskie/jc"
)

type Command func(args []string)

func getCommands() map[string]Command {
	return map[string]Command{
		"jc": jc.Main,
		"gtp": gtp.Main,

		// system commands
		"list":      list,
		"version":   showVersion,
		"-V":        showVersion,
		"--version": showVersion,
		"help":      help,
		"-h":        help,
		"--help":    help,
	}
}
