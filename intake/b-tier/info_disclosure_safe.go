package main

import (
	"encoding/json"
	"net/http"
)

// B-tier PAT-INFO-01 safe — no secrets in debug response.
func infoDisclosureSafe(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"q":      r.URL.Query().Get("q"),
	})
}
