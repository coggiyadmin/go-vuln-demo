package main
import ("net/http"; "os/exec"; "regexp")
func v07(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(q) { return }
    exec.Command("grep", q, "/var/log/app.log").Run() // SK-07 hardening gate
}
