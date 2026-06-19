// Combination #9 — COMMENT / STRING × SSTI (Go). Expect 0.
package main

import "net/http"

func commentSstiHandler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	ex := "{{ " + n + " }}"
	w.Write([]byte(ex))
}
