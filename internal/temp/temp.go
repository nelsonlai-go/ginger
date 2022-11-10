package temp

import (
	"os"

	"github.com/iancoleman/strcase"
	"github.com/nelsonlai-go/args"
	"github.com/nelsonlai-go/ginger/internal/ginger_init"
	"github.com/nelsonlai-go/ginger/pkg/format"
	"github.com/nelsonlai-go/ginger/pkg/go_script"
)

const (
	TEMP_MAPPER  = "mapper"
	TEMP_SERVICE = "service"
	TEMP_REPO    = "repo"
)

var arg = args.New()

func Temp(tempType string) {
	name := arg.FlagString("name", true, "", "n")
	overwrite := arg.FlagBool("overwrite", "o")

	name = strcase.ToCamel(name)
	interfaceName := format.FirstLetterToUpper(name) + getNameSuffix(tempType)
	implName := format.FirstLetterToLower(name) + getNameSuffix(tempType)

	script := go_script.NewScript()
	script.Package = getPackageName(tempType)

	script.AddInterface(&go_script.Interface{
		Name: interfaceName,
	})

	script.AddStruct(&go_script.Struct{
		Name: implName,
	})

	script.AddFunc(&go_script.Func{
		NameWithInputs: "New" + interfaceName + "()",
		Outputs:        interfaceName,
		Body:           "\treturn &" + implName + "{}",
	})

	scriptTxt := script.String()

	savePath := getSaveDir(tempType) + "/" + strcase.ToSnake(name) + ".go"
	if !overwrite && checkDirExists(savePath) {
		panic("File already exists")
	}
	err := os.WriteFile(savePath, []byte(scriptTxt), os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func getPackageName(tempType string) string {
	switch tempType {
	case TEMP_MAPPER:
		return "mapper"
	case TEMP_SERVICE:
		return "service"
	case TEMP_REPO:
		return "repo"
	default:
		panic("Invalid temp type")
	}
}

func getNameSuffix(tempType string) string {
	switch tempType {
	case TEMP_MAPPER:
		return "Mapper"
	case TEMP_SERVICE:
		return "Service"
	case TEMP_REPO:
		return "Repo"
	default:
		panic("Invalid temp type")
	}
}

func getSaveDir(tempType string) string {
	switch tempType {
	case TEMP_MAPPER:
		return ginger_init.PKG_MAPPER_PATH
	case TEMP_SERVICE:
		return ginger_init.PKG_SERVICE_PATH
	case TEMP_REPO:
		return ginger_init.PKG_REPO_PATH
	default:
		panic("Invalid temp type")
	}
}

func checkDirExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
