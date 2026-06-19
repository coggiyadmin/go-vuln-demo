// Combination #9 — COMMENT / STRING × DESERIALIZE (Go). Expect 0.
package main

import "net/http"

func commentDeserializeHandler(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	ex := "unserialize(s)"
	w.Write([]byte(ex + s))
}
