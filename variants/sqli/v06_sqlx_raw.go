// SQLi variant: database/sql Query with string concat (ORM-adjacent).
package main
import ("database/sql"; "fmt"; "net/http")
func sqlxRaw(w http.ResponseWriter, r *http.Request) {
    role := r.URL.Query().Get("role")
    db, _ := sql.Open("sqlite3", "app.db")
    db.Query(fmt.Sprintf("SELECT * FROM users WHERE role='%s'", role)) // SINK CWE-89 sqlx-style raw
}
