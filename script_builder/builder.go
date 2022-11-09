package script_builder

type ScriptBuilder interface {
	AddImport(alias string, path string)
	AddBody(body string)
	Build(savePath string, overwrite bool)
}

type scriptBuilder struct {
	Package string
	Imports map[string]Import
	Body    string
}

func New(packageName string) ScriptBuilder {
	return &scriptBuilder{
		Package: packageName,
		Imports: make(map[string]Import),
	}
}

func (b *scriptBuilder) AddImport(alias string, path string) {
	b.Imports[alias] = Import{
		Alias: alias,
		Path:  path,
	}
}

func (b *scriptBuilder) AddBody(body string) {
	b.Body += "\n" + body + "\n"
}
