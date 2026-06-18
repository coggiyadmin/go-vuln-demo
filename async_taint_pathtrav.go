// Combination #4 — ASYNC taint × PATH TRAVERSAL (CWE-22, Go).
package main

import (
	"net/http"
	"os"
)

func atPathHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	ch := make(chan string, 1)
	go func() { ch <- name }() // taint through goroutine
	v := <-ch
	os.ReadFile("/srv/app/data/" + v) // CWE-22
}
