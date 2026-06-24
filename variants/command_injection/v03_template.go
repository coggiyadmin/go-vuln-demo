package main
import ("os/exec"; "fmt"; "net/http")
func cmdTmpl(w http.ResponseWriter, r *http.Request) {
    host := r.FormValue("host")
    exec.Command("sh", "-c", fmt.Sprintf("ping -c 3 %s", host)).Run() // SINK CWE-78
}
