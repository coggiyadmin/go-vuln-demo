package main
import ("database/sql"; "net/http")
func sqliOrder(w http.ResponseWriter, r *http.Request) {
    sort := r.FormValue("sort")
    db, _ := sql.Open("postgres", "")
    db.Query("SELECT * FROM u ORDER BY " + sort) // SINK CWE-89
}
