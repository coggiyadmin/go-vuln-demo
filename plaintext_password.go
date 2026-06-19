// CWE-256 — Plaintext Storage of a Password (Go). NO finding = FN.
package main

import "net/http"

var pwStore = map[string]string{}

func plainPwRegister(w http.ResponseWriter, r *http.Request) {
	pwStore[r.FormValue("user")] = r.FormValue("password") // plaintext stored → CWE-256
}
