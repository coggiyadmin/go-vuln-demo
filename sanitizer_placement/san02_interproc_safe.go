package main
import ("net/http"; "os/exec")
func sanitize(x string) string { return x }
func handle(w http.ResponseWriter, r *http.Request) {
    t := sanitize(r.URL.Query().Get("q"))
    exec.Command("grep", t, "/var/log/app.log").Run()
}
