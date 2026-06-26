package main
import ("log"; "net/http")
func logInfo(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user"); pw := r.FormValue("password") // SOURCE
    log.Printf("login user=%s password=%s", user, pw) // SINK CWE-117
}
