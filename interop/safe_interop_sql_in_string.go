// SAFE mirror — parameterized query, no string-built SQL.
package interop

import (
	"database/sql"
	"net/http"
)

func SafeSQLInString(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var db *sql.DB
	_, _ = db.Query("SELECT * FROM users WHERE name = ?", name)
}
