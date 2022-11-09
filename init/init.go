package init

import (
	"os/exec"

	"github.com/nelsonlai-go/ginger/dir"
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
	dir.CreateDir(pathOfInternal)
	for _, dirPath := range []string{"app", "service", "mapper", "repo"} {
		dir.CreateDirWithGitkeep(pathOfInternal + "/" + dirPath)
	}
	var pathOfModel = pathOfInternal + "/model"
	dir.CreateDir(pathOfModel)
	for _, dirPath := range []string{"entity", "request", "response", "dto"} {
		dir.CreateDirWithGitkeep(pathOfModel + "/" + dirPath)
	}
	dir.CreateDirWithGitkeep("pkg")

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
