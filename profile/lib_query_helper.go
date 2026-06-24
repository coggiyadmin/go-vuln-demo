// FP-target (elarasu cognium-dev#128/#140) — LIBRARY profile. `where` is caller-supplied, not
// an HTTP entry point; must not be sql_injection under an entry-point gate.
package profile

import "database/sql"

func ByFilter(db *sql.DB, where string) (*sql.Rows, error) {
	return db.Query("SELECT * FROM items WHERE " + where) // caller-supplied, not entry point
}
