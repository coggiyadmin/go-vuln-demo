// Combination #4 — ASYNC TAINT × DESERIALIZE (CWE-502, Go). Taint carried through a
// goroutine + channel, then reaches the sink. NO finding = FALSE NEGATIVE.
package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
)

func atDeHandler(w http.ResponseWriter, r *http.Request) {
	raw := r.FormValue("s")
	ch := make(chan string, 1)
	go func() { ch <- raw }() // taint through goroutine
	s := <-ch
	gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502
}
