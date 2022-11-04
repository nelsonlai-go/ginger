package help

import (
	"fmt"

	"github.com/nelsonlai-go/args"
)

var arg = args.New()

func Help() {
	printHeader()
	switch len(arg.Args) {
	case 1:
		generalHelp()
	case 2:
		command := arg.Args[1]
		switch command {
		case "init":
			initHelp()
		default:
			fmt.Println("Invalid command")
		}
	}
}

func printHeader() {
	fmt.Println(`
==============
| Ginger Cli |
==============`)
}

func generalHelp() {
	fmt.Println(`
~~ Help - General ~~

Development tool for ginger engine framework.

Usage:
	ginger [OPTIONS] [COMMAND]

Commands:
	ginger help - Get help for ginger cli
	ginger help [COMMAND] - Get help for a specific command

	ginger init - Initialize ginger project
	`)
}

func initHelp() {
	fmt.Println(`
~~ Help - Init ~~

Initialize the ginger project.

Usage:
	ginger init

Options:
	(No options)
	`)
}
