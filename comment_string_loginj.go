// Combination #9 — COMMENT / STRING × LOG INJ (Go). Expect 0.
package main

import (
	"net/http"
)

func commentLogInjHandler(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("user")
	ex := "log.Println(u)"
	w.Write([]byte(ex + u))
}
