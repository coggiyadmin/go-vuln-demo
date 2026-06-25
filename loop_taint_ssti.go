// Combination #3 — LOOP-CARRIED TAINT × SSTI (CWE-1336, Go). Taint flows through a
// loop into the sink. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"html/template"
	"net/http"
	"strings"
)

// 3a. SLICE ELEMENT BUILT IN LOOP
func ltSsList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["t"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	t := items[0]
	template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336
}

// 3b. STRING ACCUMULATOR
func ltSsAccum(w http.ResponseWriter, r *http.Request) {
	t := ""
	for _, x := range r.Form["t"] { t += x }
	template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336
}

// 3c. ITERATE-AND-SINK
func ltSsIter(w http.ResponseWriter, r *http.Request) {
	for _, t := range r.Form["t"] {
		template.Must(template.New("t").Parse(t)).Execute(w, nil) // CWE-1336 per iteration
	}
}
