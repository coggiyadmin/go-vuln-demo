// Combination #13 — ENCODED × SSTI (CWE-1336, Go).
package main

import (
	"encoding/base64"
	"html/template"
	"net/http"
	"net/url"
)

func encodedSstiB64Handler(w http.ResponseWriter, r *http.Request) {
	n, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	t := template.Must(template.New("t").Parse("<p>" + string(n) + "</p>"))
	t.Execute(w, nil) // CWE-1336
}

func encodedSstiURLHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := url.QueryUnescape(r.FormValue("d"))
	t := template.Must(template.New("t").Parse("<p>" + n + "</p>"))
	t.Execute(w, nil) // CWE-1336
}
