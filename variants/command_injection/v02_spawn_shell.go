package main
import ("os/exec"; "net/http")
func cmdShell(w http.ResponseWriter, r *http.Request) {
    cmd := r.FormValue("cmd")
    exec.Command(cmd).Run() // SINK CWE-78 when shell metacharacters present
}
