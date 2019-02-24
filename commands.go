package main

import (
	csvtcli "github.com/taskie/csvt/cli"
	"github.com/taskie/fwv/cli/fwv"
	"github.com/taskie/gfp/cli/gfp"
	gtpcli "github.com/taskie/gtp/cli"
	"github.com/taskie/jc/cli/jc"
	"github.com/taskie/pity/cli/pity"
	"github.com/taskie/reinc/cli/reinc"
	"github.com/taskie/rltee/cli/rltee"
)

func getCommands() map[string]func() {
	return map[string]func(){
		"jc":    jc.Main,
		"gtp":   gtpcli.Main,
		"csvt":  csvtcli.Main,
		"fwv":   fwv.Main,
		"reinc": reinc.Main,
		"pity":  pity.Main,
		"rltee": rltee.Main,
		"gfp":   gfp.Main,

		// system commands
		"list":      list,
		"-l":        list,
		"version":   showVersion,
		"-V":        showVersion,
		"--version": showVersion,
		"help":      help,
		"-h":        help,
		"--help":    help,
	}
}
