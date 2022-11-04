package main

import (
	"log"

	"github.com/nelsonlai-go/args"
	"github.com/nelsonlai-go/ginger/help"
)

var arg = args.New()

func main() {
	if len(arg.Args) == 0 {
		log.Fatalln("No command specified")
	}
	command := arg.Args[0]

	switch command {
	case "help":
		help.Help()
	default:
		log.Fatalln("Invalid command")
	}
}
