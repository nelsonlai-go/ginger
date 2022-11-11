package ginger_init

import (
	"os/exec"

	"github.com/nelsonlai-go/args"
)

var arg = args.New()

func InitGoMod() {
	mod := arg.FlagString("mod", true, "", "m")
	err := exec.Command("go", "mod", "init", mod).Run()
	if err != nil {
		panic(err)
	}
}
