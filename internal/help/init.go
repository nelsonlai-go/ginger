package help

import "fmt"

func initHelp() {
	fmt.Println(`
init
> Initialize ginger project

Usage:
  ginger init [OPTIONS]

Options:
  --mod  (-m) | required | go mod name of the go project
  --port (-p) | optional | port number of the gin server (default: 5000)
	`)
}
