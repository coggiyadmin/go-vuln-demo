// Phase-3 encode mirror — contextual encoding before sink
package main
import ("net/http"; "regexp")
func xpathSafe(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(name) {
        http.Error(w, "forbidden", 403); return
    }
}
