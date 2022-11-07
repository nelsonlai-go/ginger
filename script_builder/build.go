package script_builder

import (
	"os"
	"os/exec"
)

func (b *scriptBuilder) Build(savePath string) {
	script := "package " + b.Package + "\n\n"

	script += "import (\n"
	for imp := range b.Imports {
		script += "\t\"" + imp + "\"\n"
	}
	script += ")\n\n"

	script += b.Body

	err := os.WriteFile(savePath, []byte(script), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = exec.Command("gofmt", "-w", savePath).Run()
	if err != nil {
		panic(err)
	}
}
