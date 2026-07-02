package main
import ("fmt"; "html"; "net/http")
func skWrongContextFmtstr(w http.ResponseWriter, r *http.Request) {
    pat := html.EscapeString(r.URL.Query().Get("fmt"))
    fmt.Fprintf(w, pat, "guest")
}
