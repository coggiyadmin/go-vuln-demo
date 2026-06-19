// SAFE mirror — dangerous_function.go; a constrained numeric add replaces the shell call.
package main

import (
	"net/http"
	"strconv"
)

func safeDangerousCalc(w http.ResponseWriter, r *http.Request) {
	a, e1 := strconv.Atoi(r.URL.Query().Get("a"))
	b, e2 := strconv.Atoi(r.URL.Query().Get("b"))
	if e1 != nil || e2 != nil {
		http.Error(w, "bad input", http.StatusBadRequest)
		return
	}
	w.Write([]byte(strconv.Itoa(a + b))) // no shell / dynamic evaluation
}
