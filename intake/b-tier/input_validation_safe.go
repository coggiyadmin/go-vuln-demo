package main

import (
	"database/sql"
	"net/http"
	"strconv"
)

const maxTransfer = 1000

// B-tier PAT-INPUT-01 safe — bounded transfer amount.
func inputValidationSafe(w http.ResponseWriter, r *http.Request) {
	amount, _ := strconv.Atoi(r.URL.Query().Get("amount"))
	if amount <= 0 || amount > maxTransfer {
		http.Error(w, "bad amount", http.StatusBadRequest)
		return
	}
	toUser := r.URL.Query().Get("to")
	db, _ := sql.Open("sqlite3", "app.db")
	_, _ = db.Exec("UPDATE accounts SET balance = balance - ? WHERE user = ?", amount, toUser)
}
