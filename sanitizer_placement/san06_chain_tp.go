package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q")
    exec.Command("sh", "-c", "grep "+t+" /var/log/app.log").Run() // san06 TP
}
