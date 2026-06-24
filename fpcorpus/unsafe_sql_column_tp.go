// TP control (CWE-89, go) — column concatenated from an HTTP source with NO validation. Proves
// the engine fires on the genuine sink, so safe_sql_identifier_quote.go staying clean means the
// validation guard is credited (not an engine blind spot).
package fpcorpus

import (
	"database/sql"
	"net/http"
)

func ByColumnUnsafe(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		col := r.URL.Query().Get("col")                      // attacker-controlled, NOT validated
		db.Query("SELECT * FROM items WHERE " + col + " = 1") // SINK
	}
}
