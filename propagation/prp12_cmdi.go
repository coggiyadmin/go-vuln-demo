package main
import ("net/http"; "os/exec")
var gStore string
func handleGlobalVarcmdi(w http.ResponseWriter, r *http.Request) {
    s := r.FormValue("q") // SOURCE
    gStore = s; t := gStore // package global
    exec.Command("sh", "-c", "grep "+t).Run() // SINK
}
