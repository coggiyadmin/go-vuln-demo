package main
import ("database/sql"; "net/http")
func handleRoundtripsqli(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    s := r.FormValue("q") // SOURCE
    b, _ := json.Marshal(map[string]string{"v": s}); var m map[string]string; json.Unmarshal(b, &m); t := m["v"] // round-trip
    db.Query("SELECT * FROM u WHERE n='" + t + "'") // SINK
}
