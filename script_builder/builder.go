package script_builder

type ScriptBuilder interface {
	AddImports(imports ...string)
	AddBody(body string)
	Build(savePath string)
}

type scriptBuilder struct {
	Package string
	Imports map[string]struct{}
	Body    string
}

func New(packageName string) ScriptBuilder {
	return &scriptBuilder{
		Package: packageName,
		Imports: make(map[string]struct{}),
	}
}

func (b *scriptBuilder) AddImports(imports ...string) {
	for _, imp := range imports {
		b.Imports[imp] = struct{}{}
	}
}

func (b *scriptBuilder) AddBody(body string) {
	b.Body += "\n" + body + "\n"
}
