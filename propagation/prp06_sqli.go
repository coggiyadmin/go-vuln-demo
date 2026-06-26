package main
import ("net/http"; "os/exec")
func h(w http.ResponseWriter, r *http.Request) {
    t := r.URL.Query().Get("q") // PRP prp06
    u := t
    v := u
    exec.Command("sh", "-c", "grep "+v).Run() // SINK
}
