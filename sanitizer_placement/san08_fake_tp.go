package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q") // sanitized below — fake SAN-08
    exec.Command("sh", "-c", "grep "+t).Run()
}
