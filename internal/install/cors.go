package install

import (
	"os"

	"github.com/nelsonlai-go/ginger/pkg/go_script"
)

func InstallCors() {
	script := go_script.Parse(getMainScript())

	script.AddImport(&go_script.Import{
		Path: "github.com/nelsonlai-go/ginger-engine/cors",
	})

	imp := script.GetImport(go_script.GetModName() + "/internal/config")
	if imp == nil {
		script.AddImport(&go_script.Import{
			Path: go_script.GetModName() + "/internal/config",
		})
	}

	if script.GetFunc("corsPlugin()") != nil {
		panic("cors already installed")
	}

	script.AddFunc(&go_script.Func{
		NameWithInputs: "corsPlugin()",
		Body:           "\tcors.Config(config.CORSOption)\n\tcors.Register(engine)",
	})

	f := script.GetFunc("installPlugin()")
	if f != nil {
		f.Body += "\tcorsPlugin()\n"
	} else {
		script.AddFunc(&go_script.Func{
			NameWithInputs: "installPlugin()",
			Body:           "\tcorsPlugin()\n",
		})
	}

	configScript := getConfigScript()
	configScript.AddImport(&go_script.Import{
		Path: "github.com/gin-contrib/cors",
	})
	configScript.AddVar(&go_script.Var{
		Name: "CORSOption",
		Value: `&cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	}`,
	})

	err := os.WriteFile("main.go", []byte(script.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("internal/config/config.go", []byte(configScript.String()), os.ModePerm)
	if err != nil {
		panic(err)
	}

	go_script.TidyGoMod()
}

func getMainScript() string {
	b, err := os.ReadFile("main.go")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getConfigScript() *go_script.Script {
	err := os.MkdirAll("internal/config", os.ModePerm)
	if err != nil {
		panic(err)
	}

	var script *go_script.Script

	_, err = os.Stat("internal/config/config.go")
	if err != nil {
		script = go_script.NewScript()
		script.Package = "config"
	} else {
		b, err := os.ReadFile("internal/config/config.go")
		if err != nil {
			panic(err)
		}
		script = go_script.Parse(string(b))
	}
	return script
}
