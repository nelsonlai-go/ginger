package init

import (
	"os"
	"os/exec"

	"github.com/nelsonlai-go/ginger/main_script"
)

/*
	Project structure will be like this:

	internal/
		cron/
			.gitkeep
		handler/
			.gitkeep
		service/
			.gitkeep
		middleware/
			.gitkeep
		model/
			entity/
				.gitkeep
			request/
				.gitkeep
			response/
				.gitkeep
			dto/
				.gitkeep
		mapper/
			.gitkeep
		repo/
			.gitkeep
	pkg/
		.gitkeep
	ui/
		.gitkeep
	go.mod
	go.sum
	main.go
*/

func InitGingerProject(mod string, port string) {
	var pathOfInternal = "internal"
	createDir(pathOfInternal)
	for _, dir := range []string{"app", "service", "mapper", "repo"} {
		createDirWithGitkeep(pathOfInternal + "/" + dir)
	}
	var pathOfModel = pathOfInternal + "/model"
	createDir(pathOfModel)
	for _, dir := range []string{"entity", "request", "response", "dto"} {
		createDirWithGitkeep(pathOfModel + "/" + dir)
	}
	createDirWithGitkeep("pkg")

	initGoMod(mod)

	mainScript := main_script.New(port)
	mainScript.Build("./main.go")
}

func initGoMod(mod string) {
	err := exec.Command("go", "mod", "init", mod).Run()
	if err != nil {
		panic(err)
	}
}

func createDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func createDirWithGitkeep(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path+"/.gitkeep", []byte(""), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
