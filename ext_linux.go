package main

import (
	"github.com/taskie/pity/cli/pity"
)

func addExtraCommands() {
	Command.AddCommand(pity.Command)
}
