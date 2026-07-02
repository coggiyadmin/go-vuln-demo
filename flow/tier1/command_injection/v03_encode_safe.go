// Phase-3 encode mirror — contextual encoding before sink
package main
import ("os/exec"; "net/http")
func cmdiSafe(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    exec.Command("grep", q, "/var/log/app.log").Run() // SAFE argv
}
