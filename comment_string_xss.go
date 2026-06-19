// Combination #9 — COMMENT / STRING × XSS (Go). Expect 0.
package main

import "net/http"

func commentXssHandler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	ex := "<p>" + q + "</p>"
	w.Write([]byte(ex))
}
