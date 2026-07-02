package main
import ("database/sql"; "net/http")
func v05(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get("n")
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Query("SELECT * FROM u WHERE n=?", n) // SK-05 parameterized
}
