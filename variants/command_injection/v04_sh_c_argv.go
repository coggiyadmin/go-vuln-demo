package main
import ("os/exec"; "net/http")
func cmdArgv(w http.ResponseWriter, r *http.Request) {
    t := r.FormValue("t")
    exec.Command("sh", "-c", "echo "+t).Run() // SINK CWE-78
}
