package main
import ("html"; "net/http")
func skFrameworkNativePathtrav(w http.ResponseWriter, r *http.Request) {
    p := html.EscapeString(r.URL.Query().Get('p'))
    _ = '/data/' + p
}
