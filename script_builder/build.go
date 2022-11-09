package script_builder

import (
	"errors"
	"os"
	"os/exec"
)

func (b *scriptBuilder) Build(savePath string, overwrite bool) {
	script := "package " + b.Package + "\n\n"

	if len(b.Imports) > 0 {
		script += "import (\n"
		for _, imp := range b.Imports {
			script += "\t"
			if imp.Alias != "" {
				script += imp.Alias + " "
			}
			script += "\"" + imp.Path + "\"\n"
		}
		script += ")\n\n"
	}

	script += b.Body

	if _, err := os.Stat(savePath); err == nil {
		if !overwrite {
			panic(errors.New("file already exists, please set overwrite to true"))
		}
	}

	err := os.WriteFile(savePath, []byte(script), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = exec.Command("gofmt", "-w", savePath).Run()
	if err != nil {
		panic(err)
	}
}
