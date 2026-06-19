// XSS (CWE-79) combination cells for Go.
package main

import (
	"encoding/base64"
	"html"
	"net/http"
	"strings"
)

func xssNegHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if strings.Contains(q, "<") {
		w.Write([]byte("<p>" + q + "</p>")) // CWE-79
	}
}

func xssWrongHandler(w http.ResponseWriter, r *http.Request) {
	q := strings.ReplaceAll(r.FormValue("q"), ";", "")
	w.Write([]byte("<p>" + q + "</p>")) // CWE-79
}

func xssWrappedHandler(w http.ResponseWriter, r *http.Request) {
	q := html.EscapeString(r.FormValue("q"))
	w.Write([]byte("<p>" + q + "</p>")) // expect 0
}

func xssB64Handler(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	w.Write([]byte("<p>" + string(raw) + "</p>")) // CWE-79
}
