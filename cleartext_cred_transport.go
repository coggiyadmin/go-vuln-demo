// CWE-523 — Unprotected Transport of Credentials (Go). Creds over plain HTTP. NO finding = FN.
package main

import (
	"net/http"
	"net/url"
)

func clearCredLogin(w http.ResponseWriter, r *http.Request) {
	http.PostForm("http://auth.internal/login", url.Values{ // http:// → CWE-523
		"u": {r.FormValue("user")}, "p": {r.FormValue("password")}})
}
