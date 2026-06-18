// SAFE mirror — oop_pathtrav.go; the tainted name is reduced to its base name and
// confined to the base dir before read. Expect 0 security findings.
package main

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const safePathBase = "/srv/app/data/"

type safePathDoc struct{ name string }

func (d safePathDoc) readDirect() ([]byte, error) {
	leaf := filepath.Base(d.name) // strip any ../ components
	full := filepath.Join(safePathBase, leaf)
	if !strings.HasPrefix(full, filepath.Clean(safePathBase)) {
		return nil, errors.New("path escapes base")
	}
	return os.ReadFile(full)
}

func safePathTravOOPHandler(w http.ResponseWriter, r *http.Request) {
	d := safePathDoc{name: r.FormValue("name")}
	d.readDirect()
}
