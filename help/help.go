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
		case "new-service":
			newServiceHelp()
		case "new-repo":
			newRepoHelp()
		case "new-mapper":
			newMapperHelp()
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
  new-service    - create new template file
  new-repo       - create new template file
  new-mapper     - create new template file
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

func newServiceHelp() {
	fmt.Println(`
Help - new-service

> Create new service template file.

Usage:
  ginger new-service [OPTIONS] - create new service template file

Options:
  --name (-n) | required | name of the service
	`)
}

func newRepoHelp() {
	fmt.Println(`
Help - new-repo

> Create new repo template file.

Usage:
  ginger new-repo [OPTIONS] - create new repo template file

Options:
  --name (-n) | required | name of the repo
	`)
}

func newMapperHelp() {
	fmt.Println(`
Help - new-mapper

> Create new mapper template file.

Usage:
  ginger new-mapper [OPTIONS] - create new mapper template file

Options:
  --name (-n) | required | name of the mapper
	`)
}
