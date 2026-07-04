package main
import ("database/sql"; "net/http")
func query(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    name := r.FormValue("name") // SOURCE
    db.Query("SELECT * FROM users WHERE name = '" + name + "'") // SINK CWE-89 database/sql concat
}
