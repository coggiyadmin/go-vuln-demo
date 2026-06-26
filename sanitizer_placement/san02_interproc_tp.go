package main
import ("net/http"; "os/exec")
func wrap(x string) string { return x }
func handle(w http.ResponseWriter, r *http.Request) {
    t := wrap(r.URL.Query().Get("q"))
    exec.Command("sh", "-c", "grep "+t).Run() // SAN-02 fake wrapper TP
}
