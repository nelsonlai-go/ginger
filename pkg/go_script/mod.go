package go_script

import (
	"os"
	"os/exec"
	"strings"
)

func GetModName() string {
	b, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	if len(lines) == 0 {
		return ""
	}

	return strings.Split(lines[0], " ")[1]
}

func TidyGoMod() {
	err := exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		panic(err)
	}
}
