package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	version  string
	revision string
)

func main() {
	commandName := filepath.Base(os.Args[0])
	args := os.Args
	if commandName == "oreutils" || commandName == "ore" {
		if len(os.Args) <= 1 {
			fmt.Fprintln(os.Stderr, "Please specify command")
			os.Exit(1)
		}
		commandName = os.Args[1]
		args = os.Args[1:]
	}

	command, ok := getCommands()[commandName]
	if !ok {
		fmt.Fprintln(os.Stderr, "Command \""+commandName+"\" is not found")
		os.Exit(1)
	}

	command(args)
}

func list(_ []string) {
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

func showVersion(_ []string) {
	fmt.Println(version)
}

const usage = `ore - my own toolchain
usage: ore COMMAND ARGS...
       ore list
detail: https://github.com/taskie/ore
`

func help(_ []string) {
	fmt.Print(usage)
}
