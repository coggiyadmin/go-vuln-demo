package main

import "net/http"

var trustedSessionSafe = map[string]map[string]string{}

func v01TrustSafe(w http.ResponseWriter, r *http.Request) {
	sid := r.Header.Get("X-Session-Id")
	if sid == "" {
		return
	}
	role := r.URL.Query().Get("role")
	if role != "user" && role != "viewer" && role != "admin" {
		return
	}
	if trustedSessionSafe[sid] == nil {
		trustedSessionSafe[sid] = map[string]string{}
	}
	trustedSessionSafe[sid]["role"] = role
}
