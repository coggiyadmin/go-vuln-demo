package main
import ("database/sql"; "net/http")
func handleClosuresqli(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    s := r.FormValue("q") // SOURCE
    f := func() string { return s }; t := f() // closure capture
    db.Query("SELECT * FROM u WHERE n='" + t + "'") // SINK
}
