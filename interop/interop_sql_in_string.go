// IL-1 polyglot — Go → SQL DSL (CWE-89).
package interop

import (
	"database/sql"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

func SQLInString(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // SOURCE
	// SINK (CWE-89): SQL guest-language assembled in a Go string literal.
	q := "SELECT * FROM users WHERE name = '" + name + "'"
	var db *sql.DB
	_, _ = db.Query(q)
}
