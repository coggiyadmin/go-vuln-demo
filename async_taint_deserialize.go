// Combination #4 — ASYNC taint × DESERIALIZE (CWE-502, Go).
package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
)

func atDeserializeHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	ch := make(chan string, 1)
	go func() { ch <- uid }()
	v := <-ch
	gob.NewDecoder(bytes.NewBufferString(uid)).Decode(new(interface{})) // CWE-502
}
