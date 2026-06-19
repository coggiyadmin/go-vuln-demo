// SAFE mirror — toctou.go; open atomically with exclusive create (no separate check).
package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func safeToctouWrite(w http.ResponseWriter, r *http.Request) {
	leaf := filepath.Base(r.URL.Query().Get("f"))
	target := filepath.Join(toctouBase, leaf)
	f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600) // atomic, no TOCTOU
	if err != nil {
		http.Error(w, "exists", http.StatusConflict)
		return
	}
	defer f.Close()
	f.WriteString(r.URL.Query().Get("d"))
}
