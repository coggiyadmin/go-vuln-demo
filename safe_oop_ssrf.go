// SAFE mirror — oop_ssrf.go; host validated against an allowlist before fetch.
// Expect 0 security findings.
package main

import (
	"errors"
	"net/http"
	"net/url"
)

var ssrfAllowed = map[string]bool{"api.internal.example.com": true}

type safeSsrfFetcher struct{ url string }

func (f safeSsrfFetcher) fetchDirect() error {
	u, err := url.Parse(f.url)
	if err != nil || !ssrfAllowed[u.Hostname()] { // allowlist enforced before sink
		return errors.New("host not allowed")
	}
	http.Get(f.url)
	return nil
}

func safeSsrfOOPHandler(w http.ResponseWriter, r *http.Request) {
	f := safeSsrfFetcher{url: r.FormValue("url")}
	_ = f.fetchDirect()
}
