package main
import ("database/sql"; "net/http")
func graphql(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id") // SOURCE CH-01
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Exec("SELECT * FROM u WHERE id='" + id + "'")
}
