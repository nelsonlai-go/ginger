package script_builder

import (
	"os"
	"strings"
)

func ReadFromScript(path string) ScriptBuilder {
	lines := readScriptAsLines(path)

	scriptCannotBeEmpty(lines)

	packageName := packageNameFromLines(lines)
	builder := New(packageName)

	imports := importsFromLines(lines)
	for _, imp := range imports {
		builder.AddImport(imp.Alias, imp.Path)
	}

	end := endOfImports(lines)
	builder.AddBody(strings.Join(lines[end:], "\n"))

	return builder
}

func readScriptAsLines(path string) []string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}

func scriptCannotBeEmpty(lines []string) {
	if len(lines) == 0 {
		panic("script is empty")
	}
}

func packageNameFromLines(lines []string) string {
	if !strings.Contains(lines[0], "package") {
		panic("package name not found")
	}
	return strings.Split(lines[0], " ")[1]
}

func importsFromLines(lines []string) []Import {
	var imports []Import
	for _, line := range lines {
		if strings.HasPrefix(line, "import") && !strings.HasPrefix(line, "import (") {
			parts := strings.Split(line, " ")
			imp := Import{}
			if len(parts) == 2 {
				imp.Path = strings.Trim(parts[1], "\"")
			} else if len(parts) == 3 {
				imp.Alias = parts[1]
				imp.Path = strings.Trim(parts[2], "\"")
			} else {
				imports = append(imports, imp)
			}
		}
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "import (") {
			for _, line := range lines {
				if strings.HasPrefix(line, ")") {
					break
				}
				parts := strings.Split(strings.Trim(line, "\t"), " ")
				imp := Import{}
				if len(parts) == 2 {
					imp.Path = strings.Trim(parts[1], "\"")
				} else if len(parts) == 3 {
					imp.Alias = parts[1]
					imp.Path = strings.Trim(parts[2], "\"")
				} else {
					imports = append(imports, imp)
				}
			}
		}
	}

	return imports
}

func endOfImports(lines []string) int {
	for i, line := range lines {
		if strings.HasPrefix(line, "import") && !strings.HasPrefix(line, "import (") {
			for _, line := range lines[i:] {
				if !strings.HasPrefix(line, "import") ||
					!strings.HasPrefix(line, "import (") ||
					!strings.HasPrefix(line, ")") ||
					!strings.HasPrefix(line, "\t") ||
					!strings.HasPrefix(line, "\n") {
					return i
				}
			}
		} else if strings.HasPrefix(line, "import (") {
			for _, line := range lines[i:] {
				if strings.HasPrefix(line, ")") {
					return i
				}
			}
		}
	}
	return 0
}
