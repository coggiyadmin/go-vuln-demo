package main
import ("database/sql"; "net/http"; "regexp")
func v07(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get("n")
    if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(n) { return }
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Query("SELECT * FROM u WHERE n=?", n)
}
