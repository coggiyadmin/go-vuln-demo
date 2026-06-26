package main
import ("net/http"; "os/exec")
func sink(v string) { exec.Command("sh", "-c", "grep "+v).Run() }
func handle(w http.ResponseWriter, r *http.Request) {
    sink(r.URL.Query().Get("q")) // SAN-07 callee TP
}
