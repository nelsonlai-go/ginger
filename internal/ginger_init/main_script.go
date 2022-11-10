package ginger_init

import (
	"fmt"
	"os"

	"github.com/nelsonlai-go/ginger/pkg/go_script"
)

func BuildMainScript() {
	goScript := go_script.NewScript()

	goScript.Package = "main"

	goScript.AddImport(&go_script.Import{
		Path: "os",
	})
	goScript.AddImport(&go_script.Import{
		Path: "github.com/nelsonlai-go/ginger-engine/ginger",
	})

	goScript.AddVar(&go_script.Var{
		Name: "engine",
		Type: "ginger.Ginger",
	})

	goScript.AddFunc(&go_script.Func{
		NameWithInputs: "main()",
		Body: `	preScript()
	installPlugin()
	postScript()`,
	})

	goScript.AddFunc(&go_script.Func{
		NameWithInputs: "preScript()",
		Body:           `	engine = ginger.New()`,
	})

	goScript.AddFunc(&go_script.Func{
		NameWithInputs: "installPlugin()",
		Body:           ``,
	})

	port := arg.FlagString("port", false, "5000", "p")

	goScript.AddFunc(&go_script.Func{
		NameWithInputs: "postScript()",
		Body: fmt.Sprintf(`	selectHost := func() string {
		mode := os.Getenv("GIN_MODE")
		if mode == "release" {
			return "0.0.0.0"
		}
		return "127.0.0.1"
	}
	engine.Run(selectHost() + ":%s")`, port),
	})

	scriptTxt := goScript.String()
	err := os.WriteFile("main.go", []byte(scriptTxt), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
