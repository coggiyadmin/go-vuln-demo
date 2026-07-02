package main
import ("html"; "net/http")
func skFrameworkNativeSsrf(w http.ResponseWriter, r *http.Request) {
    url := html.EscapeString(r.URL.Query().Get('url'))
    http.Get(url)
}
