package main
import ("net/http"; "os/exec")
func query(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q") // SOURCE SRC-01
    exec.Command("sh", "-c", "grep "+q).Run() // SINK CWE-78
}
