package main
import ("net/http"; "os/exec"; "strings")
func handle(w http.ResponseWriter, r *http.Request) {
    t := strings.ReplaceAll(r.URL.Query().Get("q"), "<", "")
    exec.Command("sh", "-c", "grep "+t).Run() // SAN-05 partial strip TP
}
