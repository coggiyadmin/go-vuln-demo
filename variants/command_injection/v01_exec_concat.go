package main
import ("os/exec"; "net/http")
func cmdExec(w http.ResponseWriter, r *http.Request) {
    q := r.FormValue("q")
    exec.Command("sh", "-c", "grep "+q).Run() // SINK CWE-78
}
