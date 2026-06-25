// Combination #2 — PATH-SENSITIVITY × LOGINJ (CWE-117, Go). Each handler is a REAL
// LOGINJ on at least one path. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"log"
	"net/http"
)

// 2a. NEGATED GUARD
func psLoNeg(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "safe" { /* covers only literal */ } else { log.Printf("login user=%s", user) } // CWE-117
}

// 2b. ONE-BRANCH CONSTRAINT
func psLoOneBranch(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if false { user = "safe_literal" } // dead branch
	log.Printf("login user=%s", user) // CWE-117
}

// 2c. EARLY-RETURN GUARD that does NOT cover the sink path
func psLoEarly(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	if user == "" { return }
	log.Printf("login user=%s", user) // CWE-117
}
