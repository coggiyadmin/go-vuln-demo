package main
import ("net/http")
func skFakeSsrf(w http.ResponseWriter, r *http.Request) {
    u := r.URL.Query().Get("url")
    http.Get("https://api.example.com/" + u)
}
