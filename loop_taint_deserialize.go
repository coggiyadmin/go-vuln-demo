// Combination #3 — LOOP-CARRIED TAINT × DESERIALIZE (CWE-502, Go). Taint flows through a
// loop into the sink. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"bytes"
	"encoding/gob"
	"net/http"
)

// 3a. SLICE ELEMENT BUILT IN LOOP
func ltDeList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["s"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	s := items[0]
	gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502
}

// 3b. STRING ACCUMULATOR
func ltDeAccum(w http.ResponseWriter, r *http.Request) {
	s := ""
	for _, x := range r.Form["s"] { s += x }
	gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502
}

// 3c. ITERATE-AND-SINK
func ltDeIter(w http.ResponseWriter, r *http.Request) {
	for _, s := range r.Form["s"] {
		gob.NewDecoder(bytes.NewBufferString(s)).Decode(new(interface{})) // CWE-502 per iteration
	}
}
