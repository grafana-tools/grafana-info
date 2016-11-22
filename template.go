package main

import (
	"text/template"
)

var (
	dashboardDisplay = template.Must(template.New("db").Parse(
		`=== Dashboard <{{.ID}}> "{{.Title}}" ===
{{if .Tags}}Tags: {{range $k,$v := .Tags}}{{printf "[%s] " $v}}{{end}}
{{end}}{{if .Templating.List}}Templating Vars: {{range $k,$v := .Templating.List}}{{printf "[%s] " $v.Name}}{{end}}
{{end}}{{range $rowI,$rowV := .Rows}}{{printf "--- Row %s ---" $rowV.Title}}
  {{if $rowV.Panels}}|{{end}}{{range $panelI,$panelV := $rowV.Panels}}{{printf " %-20.20s " $panelV.Title}}|{{end}}
  {{if $rowV.Panels}}|{{end}}{{range $panelI,$panelV := $rowV.Panels}}{{printf " %-20.20s " $panelV.Type}}|{{end}}
{{end}}
`))
	datasourceDisplay = template.Must(template.New("ds").Parse(
		`=== Datasource "{{.Name}} ==="

`))
)
