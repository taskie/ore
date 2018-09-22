package main

import (
	csvtcli "github.com/taskie/csvt/cli"
	fwvcli "github.com/taskie/fwv/cli"
	gtpcli "github.com/taskie/gtp/cli"
	jccli "github.com/taskie/jc/cli"
	leveletcli "github.com/taskie/levelet/cli"
)

type Command func()

func getCommands() map[string]Command {
	return map[string]Command{
		"jc":      jccli.Main,
		"gtp":     gtpcli.Main,
		"csvt":    csvtcli.Main,
		"fwv":     fwvcli.Main,
		"levelet": leveletcli.Main,

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
