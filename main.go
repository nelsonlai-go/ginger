package main

import (
	"log"

	"github.com/nelsonlai-go/args"
	"github.com/nelsonlai-go/ginger/internal/ginger_init"
	"github.com/nelsonlai-go/ginger/internal/help"
	"github.com/nelsonlai-go/ginger/internal/install"
	"github.com/nelsonlai-go/ginger/internal/temp"
	"github.com/nelsonlai-go/ginger/pkg/go_script"
)

var arg = args.New()

const VAR = "VAR"

const (
	IMPORT  = "import"
	PACKAGE = "package"
)

type Test struct {
	// @ginger:pre-script
	Time string `json:"time"`
	Name string `json:"name"`
}

type ITest interface {
	// @ginger:pre-script
	Build(savePath string) string
}

func main() {
	if len(arg.Args) == 0 {
		log.Fatalln("No command specified")
	}
	command := arg.Args[0]

	switch command {
	case "help":
		help.Help()
	case "init":
		ginger_init.InitGoMod()
		ginger_init.BuildProjectStructure()
		ginger_init.BuildMainScript()
		go_script.TidyGoMod()
	case "temp":
		t := arg.FlagString("type", true, "", "t")
		switch t {
		case "mapper":
			temp.Temp(temp.TEMP_MAPPER)
		case "repo":
			temp.Temp(temp.TEMP_REPO)
		case "service":
			temp.Temp(temp.TEMP_SERVICE)
		default:
			log.Fatalln("Invalid type")
		}
	case "install":
		pluginName := arg.FlagString("plugin", true, "", "p")
		switch pluginName {
		case "cors":
			install.InstallCors()
		case "ratelimit":
			install.InstallRateLimit()
		default:
			log.Fatalln("Invalid plugin")
		}
	default:
		log.Fatalln("Invalid command")
	}
}
