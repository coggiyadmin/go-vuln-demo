package main
import ("database/sql"; "encoding/json"; "net/http")
func deser(w http.ResponseWriter, r *http.Request) {
    var obj map[string]string
    json.NewDecoder(r.Body).Decode(&obj) // SOURCE SRC-11
    q := obj["q"]
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Exec("SELECT * FROM t WHERE n='" + q + "'")
}
