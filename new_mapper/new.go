package new_mapper

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/nelsonlai-go/ginger/format"
	"github.com/nelsonlai-go/ginger/script_builder"
)

func NewMapper(name string, overwrite bool) {
	script := script_builder.New("mapper")

	name = strcase.ToCamel(name)
	interfaceName := format.FirstLetterToUpper(name) + "Mapper"
	implName := format.FirstLetterToLower(name) + "Mapper"

	script.AddBody(fmt.Sprintf(`
type %s interface {}

type %s struct {}

func New%s() %s {
	return &%s{}
}
`, interfaceName, implName, interfaceName, interfaceName, implName))

	script.Build(fmt.Sprintf("internal/mapper/%s.go", strcase.ToSnake(name)), overwrite)
}
