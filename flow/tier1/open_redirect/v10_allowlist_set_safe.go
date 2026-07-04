// v10 allowlist-set — prefix slice gate before redirect.
package main

import ("net/http"; "strings")

var allowedPrefixes = []string{"https://app.example.com/", "https://cdn.example.com/"}

func openRedirect(w http.ResponseWriter, r *http.Request) {
	nxt := r.FormValue("next")
	for _, p := range allowedPrefixes {
		if strings.HasPrefix(nxt, p) {
			http.Redirect(w, r, nxt, http.StatusFound)
		return
		}
	}
	http.Redirect(w, r, "/home", http.StatusFound)
}
