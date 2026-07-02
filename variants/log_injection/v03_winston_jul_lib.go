package main
import ("log"; "net/http")
func winstonStyle(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    log.Printf("user=%s", user) // SINK CWE-117
}
