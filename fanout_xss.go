// Combination #11 — FAN-OUT × XSS (CWE-79, Go).
package main

import "net/http"

func fanoutXssHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	w.Write([]byte("<p>" + q + "</p>"))   // CWE-79
	w.Write([]byte("<h1>" + q + "</h1>"))  // CWE-79
	w.Write([]byte("<span>" + q + "</span>")) // CWE-79
}
