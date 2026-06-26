package main
import ("database/sql"; "fmt"; "net/http"; _ "github.com/mattn/go-sqlite3")
func sqli(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name") // SOURCE
    db, _ := sql.Open("sqlite3", "app.db")
    db.Query("SELECT * FROM users WHERE name='" + name + "'") // SINK CWE-89
    fmt.Fprint(w, "ok")
}
