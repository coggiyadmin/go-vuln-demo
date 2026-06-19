// SAFE mirror — path_equivalence.go; canonicalize, then enforce extension + base prefix.
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const safePeBase = "/srv/app/data/"

func safePathEquivGet(w http.ResponseWriter, r *http.Request) {
	leaf := filepath.Base(r.FormValue("name"))
	full := filepath.Join(safePeBase, leaf) // canonical form
	if !strings.HasPrefix(full, filepath.Clean(safePeBase)) || !strings.HasSuffix(full, ".txt") {
		return
	}
	os.ReadFile(full)
}
