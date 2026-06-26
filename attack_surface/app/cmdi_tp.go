package main
import ("net/http"; "os/exec")
func App(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    exec.Command("sh", "-c", "grep "+q).Run()
}
