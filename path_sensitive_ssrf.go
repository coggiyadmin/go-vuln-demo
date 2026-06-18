// Combination #2 — PATH-SENSITIVITY × SSRF (CWE-918, Go). Each handler is a REAL SSRF
// on at least one path. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"net/http"
	"strings"
)

// 2a. NEGATED GUARD — tainted URL fetched in the failure branch
func psSsrfNeg(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if !strings.HasPrefix(url, "https://api.internal/") {
		http.Get(url) // fetched anyway -> CWE-918
	}
}

// 2b. ONE-BRANCH CONSTRAINT — else path leaves the URL unchecked
func psSsrfOneBranch(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if r.FormValue("strict") != "" {
		url = "https://api.internal/" + url[strings.LastIndex(url, "/")+1:]
	}
	http.Get(url) // else path tainted -> CWE-918
}

// 2c. EARLY-RETURN GUARD that does not cover the sink path
func psSsrfEarly(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		return
	}
	http.Get(url) // CWE-918
}
