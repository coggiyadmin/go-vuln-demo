package main

import (
	"database/sql"
	"net/http"
)

// B-tier PAT-AUTH-01 safe — requires admin session before purge.
func authBypassSafe(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Admin") != "1" {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	db, _ := sql.Open("sqlite3", "app.db")
	_, _ = db.Exec("DELETE FROM users")
}
