// Combination #4 — ASYNC TAINT × SSTI (CWE-1336, Go). Taint carried through a
// goroutine + channel, then reaches the sink. NO finding = FALSE NEGATIVE.
package main

import (
	"html/template"
	"net/http"
	"strings"
)

func atSsHandler(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("t")
	ch := make(chan string, 1)
	go func() { ch <- raw }() // taint through goroutine
	t := <-ch
	template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336
}
