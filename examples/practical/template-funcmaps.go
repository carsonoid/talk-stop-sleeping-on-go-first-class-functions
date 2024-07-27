package main

import (
	"bytes"
	"strings"
	"text/template"
)

// START TEMPLATE OMIT
var templateText = `
{{- if index .featureFlags "header" }}{{ .header }}{{ end }}
{{ if index .featureFlags "caps" }}{{ toUpper .message }}{{- else }}{{ .message }}{{- end }}
{{ if index .featureFlags "footer" }}{{ .footer }}{{ end -}}
`

// END TEMPLATE OMIT

// START OMIT
var featureFlags = map[string]bool{
	"header": true,
	"caps":   true,
	"footer": true,
}

func main() {
	funcMap := template.FuncMap{"toUpper": strings.ToUpper}

	t := template.Must(template.New("t").Funcs(funcMap).Parse(templateText))
	buf := &bytes.Buffer{}
	data := map[string]interface{}{
		"featureFlags": featureFlags,
		"header":       "--- HEADER ---",
		"message":      "Hello, World!",
		"footer":       "--- FOOTER ---",
	}
	err := t.Execute(buf, data)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}

// END OMIT
