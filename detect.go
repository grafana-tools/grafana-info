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
	"github.com/grafov/autograf/grafana"
	istty "github.com/mattn/go-isatty"

	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		key        string
		db         grafana.Board
		ds         grafana.Datasource
		err        error
		isTerminal = istty.IsTerminal(os.Stdout.Fd())
	)
	flag.StringVar(&key, "key", "", "API key of Grafana server")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		// Read from file(s).
	} else {
		// Read from stdin.
		s := bufio.NewScanner(os.Stdin)
		s.Split(scanJSONLines)
		for s.Scan() {
			if err = json.Unmarshal(s.Bytes(), &ds); err == nil {
				if ds.Name != "" && ds.URL != "" {
					outputDatasource(ds, isTerminal)
					continue
				}
			}
			if err = json.Unmarshal(s.Bytes(), &db); err == nil {
				outputDashboard(db, isTerminal)
				continue
			}
			fmt.Fprintln(os.Stderr, "unknown data")
		}
	}
}

func outputDatasource(ds grafana.Datasource, isTerminal bool) {
	if isTerminal {
		datasourceDisplay.Execute(os.Stdout, ds)
	} else {
		// kv
	}
}

func outputDashboard(db grafana.Board, isTerminal bool) {
	if isTerminal {
		dashboardDisplay.Execute(os.Stdout, db)
	} else {
		// kv
	}
}

func scanJSONLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("}{")); i >= 0 {
		// We have probably a full JSON object followed by another object
		// that came obviously from another file.
		return i + 1, data[0 : i+1], nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
