package install

import (
	"os"

	"github.com/nelsonlai-go/ginger/pkg/go_script"
)

func InstallRateLimit() {
	script := go_script.Parse(getMainScript())

	script.AddImport(&go_script.Import{
		Path: "github.com/nelsonlai-go/ginger-engine/ratelimit",
	})

	imp := script.GetImport(go_script.GetModName() + "/internal/config")
	if imp == nil {
		script.AddImport(&go_script.Import{
			Path: go_script.GetModName() + "/internal/config",
		})
	}

	if script.GetFunc("rateLimitPlugin()") != nil {
		panic("rate limit already installed")
	}

	script.AddFunc(&go_script.Func{
		NameWithInputs: "rateLimitPlugin()",
		Body:           "\tratelimit.Config(config.RateLimitOption)\n\tratelimit.Register(engine)",
	})

	f := script.GetFunc("installPlugin()")
	if f != nil {
		f.Body += "\trateLimitPlugin()\n"
	} else {
		script.AddFunc(&go_script.Func{
			NameWithInputs: "installPlugin()",
			Body:           "\trateLimitPlugin()\n",
		})
	}

	configScript := getConfigScript()
	timeImp := configScript.GetImport("time")
	if timeImp == nil {
		configScript.AddImport(&go_script.Import{
			Path: "time",
		})
	}

	configScript.AddImport(&go_script.Import{
		Path: "github.com/nelsonlai-go/ginger-engine/ratelimit",
	})

	configScript.AddVar(&go_script.Var{
		Name: "RateLimitOption",
		Value: `&ratelimit.ConfigOption{
		GeneralRateLimit:         int64(60*60),
		GeneralRateLimitDuration: time.Hour,
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
