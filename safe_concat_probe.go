// SAFE mirror — concat_probe.go hardened paths (#53).
package main

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const uploadRoot = "/var/app/uploads"

func safePing(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("ping", "-c", "3", "--", host).Output()
	w.Write(out)
}

func safeRead(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	clean := filepath.Clean(filepath.Join(uploadRoot, name))
	if !strings.HasPrefix(clean, uploadRoot+string(os.PathSeparator)) {
		http.Error(w, "forbidden", 403)
		return
	}
	data, err := os.ReadFile(clean)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.Write(data)
}
