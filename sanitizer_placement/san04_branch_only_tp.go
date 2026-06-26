package main
import ("net/http"; "os/exec")
func handle(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q")
    if t == "admin" { t = "safe" }
    exec.Command("sh", "-c", "grep "+t).Run() // SAN-04 branch only TP
}
