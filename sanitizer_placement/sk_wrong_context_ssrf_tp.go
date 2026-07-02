package main
import ("html"; "net/http")
func skWrongContextSsrf(w http.ResponseWriter, r *http.Request) {
    u := html.EscapeString(r.URL.Query().Get("url"))
    http.Get("https://api.example.com/" + u)
}
