// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("html"; "fmt"; "net/http")
func xssSafe(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    fmt.Fprintf(w, "<h1>%s</h1>", html.EscapeString(q))
}
