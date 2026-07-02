package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

// B-tier PAT-INPUT-01 — unbounded transfer amount (CWE-20).
func inputValidationTp(w http.ResponseWriter, r *http.Request) {
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount")) // SINK CWE-20 — no bounds
	toUser := r.URL.Query().Get("to")
	db, _ := sql.Open("sqlite3", "app.db")
	_, _ = db.Exec("UPDATE accounts SET balance = balance - ? WHERE user = ?", amount, toUser)
}
