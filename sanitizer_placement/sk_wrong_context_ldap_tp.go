package main
import ("html"; "net/http")
func skWrongContextLdap(w http.ResponseWriter, r *http.Request) {
    uid := html.EscapeString(r.URL.Query().Get("uid"))
    _ = "(uid=" + uid + ")"
}
