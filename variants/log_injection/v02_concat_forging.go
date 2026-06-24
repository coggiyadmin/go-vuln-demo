package main
import ("log"; "net/http")
func logForge(w http.ResponseWriter, r *http.Request) {
    msg := r.FormValue("msg")
    log.Println("event: " + msg) // SINK CWE-117
}
