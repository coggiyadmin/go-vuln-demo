package main
import ("database/sql"; "net/http"; _ "github.com/mattn/go-sqlite3")
func sqliConcat(w http.ResponseWriter, r *http.Request) {
    id := r.FormValue("id")
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Query("SELECT * FROM u WHERE id=" + id) // SINK CWE-89
}
