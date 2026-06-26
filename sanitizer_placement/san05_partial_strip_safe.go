package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q")
    exec.Command("grep", t, "/var/log/app.log").Run() // san05 safe argv
}
