package main
import ("html"; "net/http")
func skFrameworkNativeLdap(w http.ResponseWriter, r *http.Request) {
    uid := html.EscapeString(r.URL.Query().Get("uid"))
    _ = "(uid=" + uid + ")"
}
