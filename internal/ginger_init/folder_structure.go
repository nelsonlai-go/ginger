package ginger_init

import "os"

const (
	PKG_INTERNAL_PATH = "internal"
	PKG_APP_PATH      = PKG_INTERNAL_PATH + "/" + "app"
	PKG_SERVICE_PATH  = PKG_INTERNAL_PATH + "/" + "service"
	PKG_MAPPER_PATH   = PKG_INTERNAL_PATH + "/" + "mapper"
	PKG_REPO_PATH     = PKG_INTERNAL_PATH + "/" + "repo"
	PKG_MODEL_PATH    = PKG_INTERNAL_PATH + "/" + "model"
	PKG_ENTITY_PATH   = PKG_MODEL_PATH + "/" + "entity"
	PKG_REQUEST_PATH  = PKG_MODEL_PATH + "/" + "request"
	PKG_RESPONSE_PATH = PKG_MODEL_PATH + "/" + "response"
	PKG_DTO_PATH      = PKG_MODEL_PATH + "/" + "dto"
	PKG_PKG_PATH      = "pkg"
)

func BuildProjectStructure() {
	for _, dirPath := range []string{
		PKG_INTERNAL_PATH,
		PKG_APP_PATH,
		PKG_SERVICE_PATH,
		PKG_MAPPER_PATH,
		PKG_REPO_PATH,
		PKG_MODEL_PATH,
		PKG_ENTITY_PATH,
		PKG_REQUEST_PATH,
		PKG_RESPONSE_PATH,
		PKG_DTO_PATH,
		PKG_PKG_PATH,
	} {
		_createDirWithGitkeep(dirPath)
	}
}

func _createDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func _createDirWithGitkeep(path string) {
	_createDir(path)
	err := os.WriteFile(path+"/.gitkeep", []byte(""), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
