package dir

import "os"

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func CreateDirWithGitkeep(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path+"/.gitkeep", []byte(""), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
