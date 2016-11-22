package main

/*
   Displays on terminal brief info about Grafana dashboards and datasources.
   Copyright (C) 2016  Alexander I.Grafov <grafov@gmail.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.


   ॐ तारे तुत्तारे तुरे स्व
*/

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
