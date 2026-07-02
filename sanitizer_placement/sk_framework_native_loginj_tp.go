package main
import ("html"; "log"; "net/http")
func skFrameworkNativeLoginj(w http.ResponseWriter, r *http.Request) {
    user := html.EscapeString(r.URL.Query().Get("user"))
    log.Printf("login user=%s", user)
}
