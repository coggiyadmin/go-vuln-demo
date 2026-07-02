package main
import ("net/http"; "strings")
func skWrongContextCrlf(w http.ResponseWriter, r *http.Request) {
    loc := strings.ReplaceAll(r.URL.Query().Get("url"), "'", "''")
    w.Header().Set("Location", loc)
}
