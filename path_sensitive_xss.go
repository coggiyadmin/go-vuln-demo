// Combination #2 — PATH-SENSITIVITY × XSS (CWE-79, Go). Neg guard in xss_combos.go.
package main

import (
	"net/http"
	"strings"
)

func xssOneBranchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if r.FormValue("strict") != "" {
		q = strings.ReplaceAll(q, "<", "")
	}
	w.Write([]byte("<p>" + q + "</p>")) // CWE-79
}

func xssEarlyHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		w.Write([]byte("empty"))
		return
	}
	w.Write([]byte("<p>" + q + "</p>")) // CWE-79
}
