// Combination #2 — PATH-SENSITIVITY × DESERIALIZE (CWE-502, Go). Each handler is a REAL
// DESERIALIZE on at least one path. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
)

// 2a. NEGATED GUARD
func psDeNeg(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	if s == "safe" { /* covers only literal */ } else { gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) } // CWE-502
}

// 2b. ONE-BRANCH CONSTRAINT
func psDeOneBranch(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	if false { s = "safe_literal" } // dead branch
	gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502
}

// 2c. EARLY-RETURN GUARD that does NOT cover the sink path
func psDeEarly(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	if s == "" { return }
	gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502
}
