// v09 guard-in-variable — host allowlist in package map.
package main

import ("net/http"; "net/url")

var allowedHosts = map[string]bool{"app.example.com": true, "cdn.example.com": true}

func openRedirect(w http.ResponseWriter, r *http.Request) {
	nxt := r.FormValue("next")
	u, err := url.Parse(nxt)
	if err == nil && allowedHosts[u.Hostname()] {
		http.Redirect(w, r, nxt, http.StatusFound)
		return
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}
