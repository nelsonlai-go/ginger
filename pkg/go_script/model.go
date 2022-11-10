package go_script

type Script struct {
	Package    string
	Imports    []*Import
	Consts     []*Const
	Vars       []*Var
	Structs    []*Struct
	Interfaces []*Interface
	Funcs      []*Func
}

func NewScript() *Script {
	return &Script{
		Imports:    make([]*Import, 0),
		Consts:     make([]*Const, 0),
		Vars:       make([]*Var, 0),
		Structs:    make([]*Struct, 0),
		Interfaces: make([]*Interface, 0),
		Funcs:      make([]*Func, 0),
	}
}

func (s *Script) GetImport(path string) *Import {
	for _, i := range s.Imports {
		if i.Path == path {
			return i
		}
	}
	return nil
}

func (s *Script) AddImport(imp *Import) {
	s.Imports = append(s.Imports, imp)
}

func (s *Script) GetConst(name string) *Const {
	for _, c := range s.Consts {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func (s *Script) AddConst(c *Const) {
	s.Consts = append(s.Consts, c)
}

func (s *Script) GetVar(name string) *Var {
	for _, v := range s.Vars {
		if v.Name == name {
			return v
		}
	}
	return nil
}

func (s *Script) AddVar(v *Var) {
	s.Vars = append(s.Vars, v)
}

func (s *Script) GetStruct(name string) *Struct {
	for _, st := range s.Structs {
		if st.Name == name {
			return st
		}
	}
	return nil
}

func (s *Script) AddStruct(st *Struct) {
	s.Structs = append(s.Structs, st)
}

func (s *Script) GetInterface(name string) *Interface {
	for _, i := range s.Interfaces {
		if i.Name == name {
			return i
		}
	}
	return nil
}

func (s *Script) AddInterface(i *Interface) {
	s.Interfaces = append(s.Interfaces, i)
}

func (s *Script) GetFunc(nameWithInputs string) *Func {
	for _, f := range s.Funcs {
		if f.NameWithInputs == nameWithInputs {
			return f
		}
	}
	return nil
}

func (s *Script) AddFunc(f *Func) {
	s.Funcs = append(s.Funcs, f)
}

func (s *Script) String() string {
	str := "package " + s.Package + "\n\n"

	if len(s.Imports) > 0 {
		str += "import (\n"
		for _, i := range s.Imports {
			str += "\t\"" + i.String() + "\"\n"
		}
		str += ")\n\n"
	}

	if len(s.Consts) > 0 {
		str += "\nconst (\n"
		for _, c := range s.Consts {
			str += "\t" + c.String() + "\n"
		}
		str += ")\n"
	}

	if len(s.Vars) > 0 {
		str += "\nvar (\n"
		for _, v := range s.Vars {
			str += "\t" + v.String() + "\n"
		}
		str += ")\n"
	}

	for _, s := range s.Structs {
		str += "\n" + s.String() + "\n"
	}

	for _, i := range s.Interfaces {
		str += "\n" + i.String() + "\n"
	}

	for _, f := range s.Funcs {
		str += "\n" + f.String() + "\n"
	}

	return str
}

type Import struct {
	Alias string
	Path  string
}

func (i *Import) String() string {
	if i.Alias == "" {
		return i.Path
	}
	return i.Alias + " " + i.Path
}

type Var struct {
	Name  string
	Type  string
	Value string
}

func (v *Var) String() string {
	s := v.Name + " " + v.Type
	if v.Value != "" {
		s += " = " + v.Value
	}
	return s
}

type Const struct {
	Name  string
	Type  string
	Value string
}

func (c *Const) String() string {
	s := c.Name + " " + c.Type
	if c.Value != "" {
		s += " = " + c.Value
	}
	return s
}

type Struct struct {
	Name   string
	Fields []StructField
}

func (s *Struct) String() string {
	str := "type " + s.Name + " struct {\n"
	for _, f := range s.Fields {
		str += "\t" + f.String() + "\n"
	}
	str += "}"
	return str
}

type StructField struct {
	Name string
	Type string
	Tags string
}

func (s *StructField) String() string {
	return s.Name + " " + s.Type + " " + s.Tags
}

type Interface struct {
	Name    string
	Methods []InterfaceMethod
}

func (i *Interface) String() string {
	str := "type " + i.Name + " interface {\n"
	for _, m := range i.Methods {
		str += "\t" + m.String() + "\n"
	}
	str += "}"
	return str
}

type InterfaceMethod struct {
	FuncName string
	Outputs  string
}

func (i *InterfaceMethod) String() string {
	return i.FuncName + " " + i.Outputs
}

type Func struct {
	StructRef      string
	NameWithInputs string
	Outputs        string
	Body           string
}

func (f *Func) String() string {
	s := "func"
	if f.StructRef != "" {
		s += " " + f.StructRef
	}
	s += " " + f.NameWithInputs

	if f.Outputs != "" {
		s += " " + f.Outputs
	}

	s += " {\n"

	s += f.Body + "\n"
	s += "}"
	return s
}
