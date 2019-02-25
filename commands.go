package main

import (
	"github.com/taskie/csvt/cli/csvt"
	"github.com/taskie/fwv/cli/fwv"
	"github.com/taskie/gfp/cli/gfp"
	"github.com/taskie/gtp/cli/gtp"
	"github.com/taskie/jc/cli/jc"
	"github.com/taskie/pity/cli/pity"
	"github.com/taskie/reinc/cli/reinc"
	"github.com/taskie/rltee/cli/rltee"
)

func getCommands() map[string]func() {
	return map[string]func(){
		"jc":    jc.Main,
		"gtp":   gtp.Main,
		"csvt":  csvt.Main,
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
