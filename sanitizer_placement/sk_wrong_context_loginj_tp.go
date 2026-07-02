package main
import ("html"; "log"; "net/http")
func skWrongContextLoginj(w http.ResponseWriter, r *http.Request) {
    user := html.EscapeString(r.FormValue("user"))
    pw := r.FormValue("password")
    log.Printf("login user=%s password=%s", user, pw)
}
