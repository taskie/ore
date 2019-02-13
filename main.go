package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/taskie/cliplus"
)

var (
	version  string
	revision string
)

func main() {
	cmd := cliplus.NewBusyCmd("ore", cliplus.NewMapMainResolver(getCommands()))
	cmd.Main()
}

func list() {
	names := make([]string, 0, 0)
	for name := range getCommands() {
		if !strings.HasPrefix(name, "-") {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}
}

func showVersion() {
	fmt.Println(version)
}

const usage = `ore - my own toolchain
usage: ore COMMAND ARGS...
       ore list
detail: https://github.com/taskie/ore
`

func help() {
	fmt.Print(usage)
}
