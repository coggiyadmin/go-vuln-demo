package main

import "net/http"

var bTierSession = map[string]map[string]string{}

func trustBoundaryTp(w http.ResponseWriter, r *http.Request) {
	sid := r.Header.Get("X-Session-Id")
	if sid == "" {
		return
	}
	bTierSession[sid] = map[string]string{"role": r.FormValue("role")} // SINK CWE-501
}
