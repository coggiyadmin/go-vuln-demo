// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("database/sql"; "net/http"; _ "github.com/mattn/go-sqlite3")
func sqliSafe(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    db, _ := sql.Open("sqlite3", "app.db")
    db.Query("SELECT * FROM users WHERE name=?", name)
}
