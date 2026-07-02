package main
import ("net/http"; "strings")
func skWrongContextCodeinj(w http.ResponseWriter, r *http.Request) {
    q := strings.ReplaceAll(r.URL.Query().Get("q"), "'", "''")
    _ = q
}
