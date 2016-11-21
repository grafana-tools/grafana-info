package main

import (
	"github.com/grafov/autograf/grafana"

	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		key string
		db  grafana.Board
		ds  grafana.Datasource
		err error
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
					datasourceDisplay.Execute(os.Stdout, ds)
					continue
				}
			}
			if err = json.Unmarshal(s.Bytes(), &db); err == nil {
				dashboardDisplay.Execute(os.Stdout, db)
				continue
			}
			fmt.Fprintln(os.Stderr, "unknown data")
		}
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
