package main

import (
	"database/sql"
	"net/http"
)

// B-tier PAT-AUTH-01 — destructive admin action without auth (CWE-306).
func authBypassTp(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "app.db")
	_, _ = db.Exec("DELETE FROM users") // SINK CWE-306
}
