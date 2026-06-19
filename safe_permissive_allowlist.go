// SAFE mirror — permissive_allowlist.go; exact host allow-list (parsed host equality).
package main

import (
	"net/http"
	"net/url"
)

var allowedFetchHosts = map[string]bool{"trusted.com": true, "www.trusted.com": true}

func safePermissiveFetch(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.Query().Get("url"))
	if err != nil || !allowedFetchHosts[u.Host] { // exact-match allow-list
		http.Error(w, "blocked", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, u.String(), http.StatusFound)
}
