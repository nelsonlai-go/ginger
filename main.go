package main

import (
	"log"

	"github.com/nelsonlai-go/args"
	"github.com/nelsonlai-go/ginger/help"
	_init "github.com/nelsonlai-go/ginger/init"
	"github.com/nelsonlai-go/ginger/new_mapper"
	"github.com/nelsonlai-go/ginger/new_repo"
	"github.com/nelsonlai-go/ginger/new_service"
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
	case "new-service":
		new_service.NewService(
			arg.FlagString("name", true, "", "n"),
		)
	case "new-repo":
		new_repo.NewRepo(
			arg.FlagString("name", true, "", "n"),
		)
	case "new-mapper":
		new_mapper.NewMapper(
			arg.FlagString("name", true, "", "n"),
		)
	default:
		log.Fatalln("Invalid command")
	}
}
