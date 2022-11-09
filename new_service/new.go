package new_service

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/nelsonlai-go/ginger/format"
	"github.com/nelsonlai-go/ginger/script_builder"
)

func NewService(name string, overwrite bool) {
	script := script_builder.New("service")

	name = strcase.ToCamel(name)
	interfaceName := format.FirstLetterToUpper(name) + "Service"
	implName := format.FirstLetterToLower(name) + "Service"

	script.AddBody(fmt.Sprintf(`
type %s interface {}

type %s struct {}

func New%s() %s {
	return &%s{}
}
`, interfaceName, implName, interfaceName, interfaceName, implName))

	script.Build(fmt.Sprintf("internal/service/%s.go", strcase.ToSnake(name)), overwrite)
}
