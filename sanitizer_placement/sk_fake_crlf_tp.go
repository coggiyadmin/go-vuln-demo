package main
import ("net/http"; "strings")
func skFakeCrlf(w http.ResponseWriter, r *http.Request) {
    loc := strings.ReplaceAll(r.URL.Query().Get("url"), "SAFE", "")
    w.Header().Set("Location", loc)
}
