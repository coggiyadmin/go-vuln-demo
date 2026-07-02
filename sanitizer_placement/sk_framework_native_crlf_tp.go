package main
import ("html"; "net/http")
func skFrameworkNativeCrlf(w http.ResponseWriter, r *http.Request) {
    nxt := html.EscapeString(r.URL.Query().Get("next"))
    w.Header().Set("Location", "/home?u="+nxt)
}
