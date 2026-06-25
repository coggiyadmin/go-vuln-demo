// Combination #4 — ASYNC TAINT × LOGINJ (CWE-117, Go). Taint carried through a
// goroutine + channel, then reaches the sink. NO finding = FALSE NEGATIVE.
package main

import (
	"log"
	"net/http"
)

func atLoHandler(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("user")
	ch := make(chan string, 1)
	go func() { ch <- raw }() // taint through goroutine
	user := <-ch
	log.Printf("login user=%s", user) // CWE-117
}
