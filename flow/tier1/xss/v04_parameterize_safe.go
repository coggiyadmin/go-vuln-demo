// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("html"; "fmt"; "net/http")
func xssSafe(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    fmt.Fprintf(w, "<h1>%s</h1>", html.EscapeString(q))
}
