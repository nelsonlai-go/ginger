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
		case "temp":
			tempHelp()
		case "install":
			installHelp()
		default:
			fmt.Println("Invalid command")
		}
	default:
		fmt.Println("Too many arguments")
	}
}

func generalHelp() {
	fmt.Println(`
> Development tool for ginger engine framework.

Usage:
  ginger [COMMAND] [OPTIONS]

Commands:
  help           - get help for ginger cli
  help [COMMAND] - get help for a specific command

  init           - initialize ginger project
  temp           - create new template file
  install        - install ginger engine plugin
	`)
}

func printHeader() {
	fmt.Println(`
==============
| Ginger Cli |
==============`)
}
