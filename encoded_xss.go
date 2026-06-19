// Combination #13 — ENCODED × XSS (CWE-79, Go).
package main

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

func encodedXssB64Handler(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	w.Write([]byte("<p>" + string(raw) + "</p>")) // CWE-79
}

func encodedXssURLHandler(w http.ResponseWriter, r *http.Request) {
	q, _ := url.QueryUnescape(r.FormValue("d"))
	w.Write([]byte("<p>" + q + "</p>")) // CWE-79
}
