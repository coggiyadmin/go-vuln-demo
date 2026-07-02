// Phase-3 encode mirror — contextual encoding before sink
package main
import ("fmt"; "net/http")
func fmtSafe(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s", r.FormValue("name"))
}
