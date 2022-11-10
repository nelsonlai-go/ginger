package go_script

import (
	"strings"
)

const (
	PACKAGE   = "PACKAGE"
	IMPORT    = "IMPORT"
	CONST     = "CONST"
	VAR       = "VAR"
	STRUCT    = "STRUCT"
	INTERFACE = "INTERFACE"
	FUNC      = "FUNC"
	SKIP      = "SKIP"
)

// Parser is a parser for go script
type parser struct {
	CurrentLine int
	ScriptLines []string
	Script      *Script
}

func Parse(script string) *Script {
	lines := strings.Split(script, "\n")
	p := &parser{
		CurrentLine: 0,
		ScriptLines: lines,
		Script:      NewScript(),
	}

	for p.CurrentLine < len(p.ScriptLines) {
		lineType := p.determineLineType(p.ScriptLines[p.CurrentLine])

		switch lineType {
		case PACKAGE:
			p.parsePackage()
		case IMPORT:
			p.parseImport()
		case CONST:
			p.parseConstOrVar(CONST)
		case VAR:
			p.parseConstOrVar(VAR)
		case STRUCT:
			p.parseStruct()
		case INTERFACE:
			p.parseInterface()
		case FUNC:
			p.parseFunc()
		default:
			p.CurrentLine++
		}
	}

	return p.Script
}

func (p *parser) Line() string {
	return p.ScriptLines[p.CurrentLine]
}

func (p *parser) OffsetLine(n int) string {
	return p.ScriptLines[p.CurrentLine+n]
}

func (p *parser) parsePackage() {
	xs := strings.Split(p.Line(), " ")
	p.Script.Package = xs[1]
	p.CurrentLine++
}

func (p *parser) parseImport() {
	isImportGroup := strings.HasPrefix(p.Line(), "import (")
	if isImportGroup {
		p.parseImportGroup()
	} else {
		p.parseImportSingle()
	}
}

func (p *parser) parseImportGroup() {
	p.CurrentLine++
	for !strings.HasPrefix(p.Line(), ")") {
		imp := p._parseImport(p.Line())
		if imp.Path != "" {
			p.Script.AddImport(&imp)
		}
		p.CurrentLine++
	}
	p.CurrentLine++
}

func (p *parser) parseImportSingle() {
	line := strings.Trim(p.Line(), "import")
	imp := p._parseImport(line)
	p.Script.AddImport(&imp)
	p.CurrentLine++
}

func (p *parser) _parseImport(line string) Import {
	line = trimLine(line)
	line = trimDoubleSpace(line)
	xs := strings.Split(line, " ")
	if len(xs) == 2 {
		return Import{
			Alias: xs[0],
			Path:  strings.Trim(xs[1], "\""),
		}
	}
	return Import{
		Alias: "",
		Path:  strings.Trim(xs[0], "\""),
	}
}

func (p *parser) parseConstOrVar(parseType string) {
	line := p.Line()
	isGroup := strings.HasPrefix(line, strings.ToLower(parseType)+" (")
	if isGroup {
		p.parseConstOrVarGroup(parseType)
	} else {
		p.parseConstOrVarSingle(parseType)
	}
	p.CurrentLine++
}

func (p *parser) parseConstOrVarGroup(parseType string) {
	p.CurrentLine++
	for !strings.HasPrefix(p.Line(), ")") {
		if parseType == CONST {
			c := p._parseConst(p.Line())
			if c.Name != "" {
				p.Script.AddConst(&c)
			}
		} else if parseType == VAR {
			v := p._parseVar(p.Line())
			if v.Name != "" {
				p.Script.AddVar(&v)
			}
		}
		p.CurrentLine++
	}
	p.CurrentLine++
}

func (p *parser) parseConstOrVarSingle(parseType string) {
	line := p.Line()
	if parseType == CONST {
		c := p._parseConst(line)
		p.Script.AddConst(&c)
	} else if parseType == VAR {
		v := p._parseVar(line)
		p.Script.AddVar(&v)
	}
	p.CurrentLine++
}

func (p *parser) _parseConst(line string) Const {

	getNameAndType := func(s string) (string, string) {
		s = trimLine(s)
		s = trimDoubleSpace(s)
		s = strings.TrimPrefix(s, "const")
		xs := strings.Split(s, " ")
		if len(xs) == 2 {
			return xs[0], xs[1]
		}
		return xs[0], ""
	}

	if strings.Contains(line, "=") {
		xs := strings.Split(line, "=")
		name, typ := getNameAndType(xs[0])
		return Const{
			Name:  name,
			Type:  typ,
			Value: strings.Trim(xs[1], " "),
		}
	}

	name, typ := getNameAndType(line)
	return Const{
		Name:  name,
		Type:  typ,
		Value: "",
	}
}

func (p *parser) _parseVar(line string) Var {

	getNameAndType := func(s string) (string, string) {
		s = trimLine(s)
		s = trimDoubleSpace(s)
		s = strings.TrimPrefix(s, "var")
		xs := strings.Split(s, " ")
		if len(xs) == 2 {
			return xs[0], xs[1]
		}
		return xs[0], ""
	}

	if strings.Contains(line, "=") {
		xs := strings.Split(line, "=")
		name, typ := getNameAndType(xs[0])
		return Var{
			Name:  name,
			Type:  typ,
			Value: strings.Trim(xs[1], " "),
		}
	}

	name, typ := getNameAndType(line)
	return Var{
		Name:  name,
		Type:  typ,
		Value: "",
	}
}

func (p *parser) parseStruct() {
	name := strings.Split(p.Line(), " ")[1]
	s := Struct{
		Name:   name,
		Fields: make([]StructField, 0),
	}
	p.CurrentLine++
	for !strings.HasPrefix(p.Line(), "}") {
		line := trimLine(p.Line())
		line = trimDoubleSpace(line)
		if isComment(line) {
			p.CurrentLine++
			continue
		}
		if line == "" {
			p.CurrentLine++
			continue
		}
		xs := strings.Split(line, " ")
		f := StructField{
			Name: xs[0],
			Type: xs[1],
		}
		if len(xs) == 3 {
			f.Tags = xs[2]
		}
		s.Fields = append(s.Fields, f)
		p.CurrentLine++
	}
	p.Script.AddStruct(&s)
	p.CurrentLine++
}

func (p *parser) parseInterface() {
	name := strings.Split(p.Line(), " ")[1]
	i := Interface{
		Name:    name,
		Methods: make([]InterfaceMethod, 0),
	}
	p.CurrentLine++
	for !strings.HasPrefix(p.Line(), "}") {
		line := trimLine(p.Line())
		line = trimDoubleSpace(line)
		if isComment(line) {
			p.CurrentLine++
			continue
		}
		if line == "" {
			p.CurrentLine++
			continue
		}
		xs := strings.Split(line, " ")
		m := InterfaceMethod{
			FuncName: xs[0],
			Outputs:  strings.Join(xs[1:], " "),
		}
		i.Methods = append(i.Methods, m)
		p.CurrentLine++
	}
	p.Script.AddInterface(&i)
	p.CurrentLine++
}

func (p *parser) parseFunc() {
	f := Func{}

	line := strings.TrimSuffix(p.Line(), "{")
	line = trimLine(line)
	line = trimDoubleSpace(line)

	xs := strings.Split(line, " ")
	if strings.HasPrefix(xs[1], "(") && strings.HasSuffix(xs[1], ")") {
		f.StructRef = xs[0]
		f.NameWithInputs = xs[2]
		f.Outputs = strings.Join(xs[3:], " ")
	} else {
		f.NameWithInputs = xs[1]
		f.Outputs = strings.Join(xs[2:], " ")
	}
	p.CurrentLine++

	body := ""
	for !strings.HasPrefix(p.Line(), "}") {
		body += p.Line() + "\n"
		p.CurrentLine++
	}
	f.Body = body

	p.Script.AddFunc(&f)
}

func (p *parser) determineLineType(line string) string {
	if strings.HasPrefix(line, "package") {
		return PACKAGE
	} else if strings.HasPrefix(line, "import") {
		return IMPORT
	} else if strings.HasPrefix(line, "type") {
		xs := strings.Split(line, " ")
		if xs[2] == "struct" {
			return STRUCT
		} else if xs[2] == "interface" {
			return INTERFACE
		}
	} else if strings.HasPrefix(line, "const") {
		return CONST
	} else if strings.HasPrefix(line, "var") {
		return VAR
	} else if strings.HasPrefix(line, "func") {
		return FUNC
	}
	return SKIP
}

func trimLine(line string) string {
	line = strings.Trim(line, " ")
	line = strings.Trim(line, "\t")
	return line
}

func trimDoubleSpace(line string) string {
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}
	return line
}

func isComment(line string) bool {
	line = trimLine(line)
	return strings.HasPrefix(line, "//")
}
