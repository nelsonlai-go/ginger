package main_script

import (
	"fmt"
	"os/exec"

	"github.com/nelsonlai-go/ginger/script_builder"
)

type MainScript interface {
	Build(savePath string)
}

type mainScript struct {
	Builder script_builder.ScriptBuilder
	Port    string
}

func New(port string) MainScript {
	return &mainScript{
		Builder: script_builder.New("main"),
		Port:    port,
	}
}

func (s *mainScript) Build(savePath string) {
	s.Builder.AddImport("", "os")
	s.Builder.AddImport("", "github.com/nelsonlai-go/ginger-engine/ginger")

	s.Builder.AddBody(s.startScript())
	s.Builder.AddBody(s.endScript())
	s.Builder.Build(savePath, true)

	err := exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		panic(err)
	}
}

func (s *mainScript) startScript() string {
	return `
func main() {
	e := ginger.New()
	`
}

func (s *mainScript) endScript() string {
	return fmt.Sprintf(`
	selectHost := func() string {
		mode := os.Getenv("GIN_MODE")
		if mode == "release" {
			return "0.0.0.0"
		}
		return "127.0.0.1"
	}
	e.Run(selectHost() + ":" + "%s")
}
	`, s.Port)
}
