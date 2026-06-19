// Combination #3 — LOOP × XSS (CWE-79, Go).
package main

import (
	"net/http"
	"strings"
)

func loopTaintXssHandler(w http.ResponseWriter, r *http.Request) {
	for _, item := range r.Form["q"] {
		w.Write([]byte("<span>" + item + "</span>")) // CWE-79
	}
	if len(r.Form["q"]) == 0 {
		for _, item := range strings.Split(r.FormValue("q"), ",") {
			w.Write([]byte("<span>" + item + "</span>")) // CWE-79
		}
	}
}
