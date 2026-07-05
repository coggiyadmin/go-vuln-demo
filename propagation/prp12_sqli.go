package main
import ("database/sql"; "net/http")
var gStore string
func handleGlobalVarsqli(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    s := r.FormValue("q") // SOURCE
    gStore = s; t := gStore // package global
    db.Query("SELECT * FROM u WHERE n='" + t + "'") // SINK
}
