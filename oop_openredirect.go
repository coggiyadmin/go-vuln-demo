// Combination #5 — OOP OBJECT FLOW × OPEN REDIRECT (CWE-601, Go). Taint stored in
// a struct field, used in a method sink. NO finding = FALSE NEGATIVE.
package main

import "net/http"

type redirNav struct{ url string }

func (n redirNav) goDirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, n.url, http.StatusFound) // open-redirect sink (CWE-601)
}

func openRedirectOOPHandler(w http.ResponseWriter, r *http.Request) {
	n := redirNav{url: r.FormValue("next")} // SOURCE -> struct field
	n.goDirect(w, r)
}
