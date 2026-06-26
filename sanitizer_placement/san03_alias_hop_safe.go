package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q")
    u := t
    exec.Command("grep", u, "/var/log/app.log").Run()
}
