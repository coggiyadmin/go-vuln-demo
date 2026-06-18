// Combination #4 — ASYNC taint × SSRF (CWE-918, Go). Taint carried through a goroutine
// + channel, then fetched. NO finding = FALSE NEGATIVE.
package main

import "net/http"

func atSsrfHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	ch := make(chan string, 1)
	go func() { ch <- url }() // taint through goroutine
	v := <-ch
	http.Get(v) // CWE-918
}
