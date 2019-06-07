package main

import (
	"html/template"
	"os"
)

func main() {
	const tpl = `
<body>

	{{.Escaped}}
	
	{{.NotEscaped}}
	
</body>
	`

	t, _ := template.New("webpage").Parse(tpl)

	data := struct {
		Escaped    string
		NotEscaped template.HTML
	}{
		Escaped:    "<b>escaped<b>",
		NotEscaped: "<b>not escaped</b>",
	}

	_ = t.Execute(os.Stdout, data)
}
