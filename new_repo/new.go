package new_repo

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/nelsonlai-go/ginger/format"
	"github.com/nelsonlai-go/ginger/script_builder"
)

func NewRepo(name string, overwrite bool) {
	script := script_builder.New("repo")

	name = strcase.ToCamel(name)
	interfaceName := format.FirstLetterToUpper(name) + "Repo"
	implName := format.FirstLetterToLower(name) + "Repo"

	script.AddBody(fmt.Sprintf(`
type %s interface {}

type %s struct {}

func New%s() %s {
	return &%s{}
}
`, interfaceName, implName, interfaceName, interfaceName, implName))

	script.Build(fmt.Sprintf("internal/repo/%s.go", strcase.ToSnake(name)), overwrite)
}
