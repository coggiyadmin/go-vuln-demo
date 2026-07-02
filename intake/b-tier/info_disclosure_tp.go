package main

import (
	"encoding/json"
	"net/http"
	"os"
)

// B-tier PAT-INFO-01 — debug handler leaks env (CWE-200).
func infoDisclosureTp(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{
		"token": os.Getenv("API_TOKEN"),
		"q":     r.URL.Query().Get("q"),
	}) // SINK CWE-200
}
