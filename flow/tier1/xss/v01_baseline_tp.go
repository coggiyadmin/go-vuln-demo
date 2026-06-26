package main
import ("fmt"; "net/http")
func xss(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q") // SOURCE
    fmt.Fprintf(w, "<h1>%s</h1>", q) // SINK CWE-79
}
