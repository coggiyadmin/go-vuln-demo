// XSS variant: reflected HTML body.
package main
import ("fmt"; "net/http")
func xssReflect(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q") // SOURCE
    fmt.Fprintf(w, "<h1>%s</h1>", q) // SINK CWE-79
}
