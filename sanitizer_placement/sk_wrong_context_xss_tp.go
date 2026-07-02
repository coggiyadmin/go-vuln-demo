package main
import ("net/http"; "strings")
func skWrongContextXss(w http.ResponseWriter, r *http.Request) {
    q := strings.ReplaceAll(r.URL.Query().Get("q"), "'", "''")
    w.Write([]byte("<p>" + q + "</p>"))
}
