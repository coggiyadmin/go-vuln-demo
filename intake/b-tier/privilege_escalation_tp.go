package main

import "net/http"

// B-tier PAT-PRIV-01 — query flag grants admin role (CWE-269).
func privilegeEscalationTp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("admin") == "1" {
		r.Header().Set("X-Role", "admin") // SINK CWE-269
	}
}
