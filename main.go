package main

import (
	"log"

	"github.com/nelsonlai-go/args"
	"github.com/nelsonlai-go/ginger/help"
	_init "github.com/nelsonlai-go/ginger/init"
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
	case "init":
		_init.InitGingerProject(
			arg.FlagString("mod", true, "", "m"),
			arg.FlagString("port", false, "5000", "p"),
		)
	default:
		log.Fatalln("Invalid command")
	}
}
