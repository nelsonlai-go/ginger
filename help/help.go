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
Help - General

> Development tool for ginger engine framework.

Usage:
  ginger [COMMAND] [OPTIONS]

Commands:
  help           - get help for ginger cli
  help [COMMAND] - get help for a specific command

  init           - initialize ginger project
	`)
}

func initHelp() {
	fmt.Println(`
Help - init

> Initialize the ginger project.

Usage:
  ginger init [OPTIONS] - initialize ginger project (golang)

Dependencies:
  - You should install go 1.16 or above

Options:
  --mod  (-m) | required | go mod name of the go project
  --port (-p) | optional | port number of the gin server (default: 5000)
	`)
}
