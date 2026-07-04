package main
import ("log"; "net/http")
func audit(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user") // SOURCE
    log.Printf("login user=%s", user) // SINK CWE-117 go log user newline
}
