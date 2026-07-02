package main
import ("net/http"; "strings")
func skFakeXss(w http.ResponseWriter, r *http.Request) {
    q := strings.ReplaceAll(r.URL.Query().Get("q"), "SANITIZE", "")
    w.Write([]byte("<p>" + q + "</p>"))
}
