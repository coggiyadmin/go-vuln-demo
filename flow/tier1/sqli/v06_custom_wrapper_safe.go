package main
import ("database/sql"; "net/http")
func v06(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get("n")
    db, _ := sql.Open("sqlite3", ":memory:")
    safeLookup(db, n)
}
func safeLookup(db *sql.DB, n string) { db.Query("SELECT * FROM u WHERE n=?", n) }
