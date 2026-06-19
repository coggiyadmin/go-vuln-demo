// SAFE mirror — external_file_path.go; basename + base-dir confinement. Expect 0.
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const extFileBase = "/srv/app/public/"

func safeExtFileRead(w http.ResponseWriter, r *http.Request) {
	leaf := filepath.Base(r.FormValue("path"))
	full := filepath.Join(extFileBase, leaf)
	if !strings.HasPrefix(full, filepath.Clean(extFileBase)) {
		return
	}
	os.ReadFile(full)
}
