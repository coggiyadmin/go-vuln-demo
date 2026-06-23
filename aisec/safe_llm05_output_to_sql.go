// SAFE mirror (OWASP LLM05) — constant parameterized query per allowed field.
package aisec

import "database/sql"

func Search(db *sql.DB, field, value string) (*sql.Rows, error) {
	switch field {
	case "name":
		return db.Query("SELECT * FROM items WHERE name = ?", value)
	case "sku":
		return db.Query("SELECT * FROM items WHERE sku = ?", value)
	}
	return nil, sql.ErrNoRows
}
