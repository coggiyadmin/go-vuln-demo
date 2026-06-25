// Combination #2 — PATH-SENSITIVITY × SSTI (CWE-1336, Go). Each handler is a REAL
// SSTI on at least one path. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"html/template"
	"net/http"
	"strings"
)

// 2a. NEGATED GUARD
func psSsNeg(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("t")
	if t == "safe" { /* covers only literal */ } else { template.Must(template.New("t").Parse(t)).Execute(w, nil) } // CWE-1336
}

// 2b. ONE-BRANCH CONSTRAINT
func psSsOneBranch(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("t")
	if false { t = "safe_literal" } // dead branch
	template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336
}

// 2c. EARLY-RETURN GUARD that does NOT cover the sink path
func psSsEarly(w http.ResponseWriter, r *http.Request) {
	t := r.FormValue("t")
	if t == "" { return }
	template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336
}
