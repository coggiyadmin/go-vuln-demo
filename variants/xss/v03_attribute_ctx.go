package main
import ("fmt"; "net/http")
func xssAttr(w http.ResponseWriter, r *http.Request) {
    u := r.URL.Query().Get("u")
    fmt.Fprintf(w, `<a href="%s">x</a>`, u) // SINK CWE-79
}
