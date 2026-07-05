package main
import ("net/http"; "os/exec")
func handleClosurecmdi(w http.ResponseWriter, r *http.Request) {
    s := r.FormValue("q") // SOURCE
    f := func() string { return s }; t := f() // closure capture
    exec.Command("sh", "-c", "grep "+t).Run() // SINK
}
