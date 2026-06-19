// CWE-93 — CRLF Injection (Go). User input with CR/LF set into a response header. FN probe.
package main

import "net/http"

func crlfTrack(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("v")        // SOURCE
	w.Header().Set("X-Track", val) // CR/LF injects/splits headers → CWE-93
}
