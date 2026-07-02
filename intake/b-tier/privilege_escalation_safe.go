package main

import "net/http"

// B-tier PAT-PRIV-01 safe — admin role only after server-side auth check.
func privilegeEscalationSafe(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Authenticated-Admin") != "1" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	r.Header().Set("X-Role", "admin")
}
