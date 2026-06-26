// Combination #3 — LOOP-CARRIED TAINT × DESERIALIZE (CWE-502, Go).
package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
)

func ltDeserializeList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["uid"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	uid := items[0]
	gob.NewDecoder(bytes.NewBufferString(uid)).Decode(new(interface{})) // CWE-502
}
