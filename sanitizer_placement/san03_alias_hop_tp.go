package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    raw := r.URL.Query().Get("q")
    _ = raw // sanitized copy ignored
    exec.Command("sh", "-c", "grep "+raw).Run() // SAN-03 alias hop TP
}
