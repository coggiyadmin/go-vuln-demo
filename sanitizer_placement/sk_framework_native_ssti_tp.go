package main
import ("html"; "net/http")
func skFrameworkNativeSsti(w http.ResponseWriter, r *http.Request) {
    uid := html.EscapeString(r.URL.Query().Get("uid"))
    _ = "SELECT * FROM u WHERE n='" + uid + "'"
}
