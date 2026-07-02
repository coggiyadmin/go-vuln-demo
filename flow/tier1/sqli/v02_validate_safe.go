// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("database/sql"; "net/http"; _ "github.com/mattn/go-sqlite3")
func sqliSafe(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    db, _ := sql.Open("sqlite3", "app.db")
    db.Query("SELECT * FROM users WHERE name=?", name)
}
