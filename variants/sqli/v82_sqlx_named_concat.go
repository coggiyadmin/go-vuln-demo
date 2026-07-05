package main
import ("net/http"; "github.com/jmoiron/sqlx")
func query(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
    name := r.FormValue("name") // SOURCE
    db.Queryx("SELECT * FROM users WHERE name = '" + name + "'") // SINK CWE-89 sqlx concat
}
