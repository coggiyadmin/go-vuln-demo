package main
import ("net/http"; "os/exec")
func v06(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    safeGrep(q) // SK-06 custom wrapper
}
func safeGrep(q string) error { return exec.Command("grep", q, "/var/log/app.log").Run() }
