// SAFE mirror — ext_control_config.go; only an allow-listed setting may be changed.
package main

import (
	"net/http"
	"os"
)

var allowedConfigKeys = map[string]bool{"locale": true, "theme": true}

func safeExtControlConfig(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k")
	if !allowedConfigKeys[key] { // allow-listed keys only
		http.Error(w, "forbidden key", http.StatusBadRequest)
		return
	}
	os.Setenv("APP_"+key, r.URL.Query().Get("v"))
}
