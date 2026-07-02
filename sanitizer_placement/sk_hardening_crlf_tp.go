package main
import ("net/http"; "strings")
func skHardeningCrlf(w http.ResponseWriter, r *http.Request) {
    loc := strings.ReplaceAll(r.URL.Query().Get("url"), "\n", "")
    w.Header().Set("Location", loc)
}
