package main

import "net/http"

var trustedSession = map[string]map[string]string{}

func v01TrustTp(w http.ResponseWriter, r *http.Request) {
	sid := r.Header.Get("X-Session-Id")
	if sid == "" {
		return
	}
	role := r.URL.Query().Get("role") // SOURCE
	if trustedSession[sid] == nil {
		trustedSession[sid] = map[string]string{}
	}
	trustedSession[sid]["role"] = role // SINK (CWE-501)
}
