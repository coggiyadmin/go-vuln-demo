// Framework-idiom benign (WS-5.2) — GORM / database/sql with placeholders that LOOK
// like query building but parameterize. ZERO security findings expected.
package main

import (
	"database/sql"
	"net/http"
)

func handleORM(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	name := r.URL.Query().Get("name")
	// Placeholder binding — `name` is a bound parameter, never concatenated.
	rows, _ := db.Query("SELECT id, email FROM users WHERE name = $1", name)
	defer rows.Close()
}
