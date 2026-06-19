// SAFE mirror — system_info_exposure.go; returns only a static health status.
package main

import "net/http"

func safeDebugInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`)) // no system/env information disclosed
}
