package main
import ("database/sql"; "net/http")
func sqliLike(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    db, _ := sql.Open("mysql", "")
    db.Query("SELECT * FROM u WHERE name LIKE '%" + name + "%'") // SINK CWE-89
}
