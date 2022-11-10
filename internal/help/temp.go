package help

func tempHelp() {
	println(`
temp
> Create new template file

Usage:
  ginger temp [OPTIONS]

Options:
  --type (-t)      | required | type of the template file (service, repo, mapper)
  --name (-n)      | required | name of the template file
  --overwrite (-o) | optional | overwrite existing template file
`)
}
