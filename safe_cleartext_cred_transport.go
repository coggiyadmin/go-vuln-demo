// SAFE mirror — cleartext_cred_transport.go; credentials sent over HTTPS. Expect 0.
package main

import (
	"net/http"
	"net/url"
)

func safeClearCredLogin(w http.ResponseWriter, r *http.Request) {
	http.PostForm("https://auth.internal/login", url.Values{
		"u": {r.FormValue("user")}, "p": {r.FormValue("password")}})
}
