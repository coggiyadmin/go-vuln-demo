// Improper Output Handling (OWASP LLM05) — model output concatenated into SQL. TP.
package aisec

import "database/sql"

func Search(db *sql.DB, clause string) (*sql.Rows, error) {
	// SINK (LLM05 -> SQLi): model-produced clause concatenated into the query
	return db.Query("SELECT * FROM items WHERE " + clause)
}
