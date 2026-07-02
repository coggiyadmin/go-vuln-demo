package main
import ("net/http"; "os/exec")
func v05(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    exec.Command("grep", q, "/var/log/app.log").Run() // SK-05 framework argv
}
