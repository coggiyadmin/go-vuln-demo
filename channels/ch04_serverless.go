package main
import ("net/http"; "os/exec")
func serverless(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q") // SOURCE CH-04
    exec.Command("sh", "-c", "grep "+q).Run()
}
