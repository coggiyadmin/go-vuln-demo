package main
import ("fmt"; "net/http"; "strings")
func skFakeFmtstr(w http.ResponseWriter, r *http.Request) {
    pat := strings.ReplaceAll(r.URL.Query().Get("fmt"), "SAFE", "")
    fmt.Fprintf(w, pat, "guest")
}
