package main

import (
	"bytes"
	"strings"
	"text/template"
)

// START TEMPLATE OMIT
var templateText = `
{{- if flag "header" }}{{ .header }}{{ end }}
{{ if flag "caps" }}{{ toUpper .message }}{{- else }}{{ .message }}{{ end }}
{{ if flag "footer" }}{{ .footer }}{{ end -}}
`

// END TEMPLATE OMIT

// START OMIT
var featureFlags = map[string]bool{
	"header": true, "caps": true, "footer": true,
}

func main() {
	funcMap := template.FuncMap{
		"toUpper": strings.ToUpper,
		"flag":    func(key string) bool { return featureFlags[key] },
	}

	t := template.Must(template.New("t").Funcs(funcMap).Parse(templateText))
	buf := &bytes.Buffer{}
	data := map[string]interface{}{
		"header":  "--- HEADER ---",
		"message": "Hello, World!",
		"footer":  "--- FOOTER ---",
	}
	err := t.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}

// END OMIT
