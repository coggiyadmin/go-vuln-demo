// Combination #3 — LOOP-CARRIED TAINT × SSRF (CWE-918, Go). Taint flows through a loop
// into the fetch sink. A handler with NO finding is a FALSE NEGATIVE.
package main

import "net/http"

// 3a. SLICE ELEMENT BUILT IN LOOP
func ltSsrfList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["u"]
	urls := make([]string, 0, len(in))
	for _, u := range in {
		urls = append(urls, u)
	}
	http.Get(urls[0]) // taint via slice element -> CWE-918
}

// 3b. STRING ACCUMULATOR
func ltSsrfAccum(w http.ResponseWriter, r *http.Request) {
	acc := "https://"
	for _, p := range r.Form["p"] {
		acc += p
	}
	http.Get(acc) // accumulated tainted URL -> CWE-918
}

// 3c. ITERATE-AND-SINK
func ltSsrfIter(w http.ResponseWriter, r *http.Request) {
	for _, u := range r.Form["u"] {
		http.Get(u) // CWE-918 per iteration
	}
}
