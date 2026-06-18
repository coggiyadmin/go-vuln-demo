// Combination #5 — OOP OBJECT FLOW × SSRF (CWE-918, Go). Taint stored in a struct
// field, fetched in a method sink. NO finding = FALSE NEGATIVE.
package main

import "net/http"

type ssrfFetcher struct{ url string }

func (f ssrfFetcher) fetchDirect() (*http.Response, error) {
	return http.Get(f.url) // SSRF sink (CWE-918)
}

func ssrfOOPHandler(w http.ResponseWriter, r *http.Request) {
	f := ssrfFetcher{url: r.FormValue("url")} // SOURCE -> struct field
	f.fetchDirect()
}
