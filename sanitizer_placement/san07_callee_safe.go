package main
import ("net/http"; "os/exec")
func sink(v string) { exec.Command("grep", v, "/var/log/app.log").Run() }
func handle(w http.ResponseWriter, r *http.Request) {
    sink(r.URL.Query().Get("q"))
}
