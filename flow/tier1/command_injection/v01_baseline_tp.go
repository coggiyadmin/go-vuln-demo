package main
import ("os/exec"; "net/http")
func cmdi(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q") // SOURCE
    exec.Command("sh", "-c", "grep "+q).Run() // SINK CWE-78
}
