// Combination #2 — PATH-SENSITIVITY × PATH TRAVERSAL (CWE-22, Go).
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const psPathBase = "/srv/app/data/"

func psPathNeg(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if strings.Contains(name, "..") || strings.Contains(name, "/") {
		os.ReadFile(psPathBase + name) // read anyway -> CWE-22
	}
}

func psPathOneBranch(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if r.FormValue("strict") != "" {
		name = filepath.Base(name)
	}
	os.ReadFile(psPathBase + name) // else path tainted -> CWE-22
}

func psPathEarly(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		return
	}
	os.ReadFile(psPathBase + name) // CWE-22
}
