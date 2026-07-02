package main
import ("log"; "net/http"; "strings")
func skHardeningLoginj(w http.ResponseWriter, r *http.Request) {
    user := strings.ReplaceAll(r.FormValue("user"), "\n", "")
    pw := strings.ReplaceAll(r.FormValue("password"), "\n", "")
    log.Printf("login user=%s password=%s", user, pw)
}
