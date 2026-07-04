package main
import ("io"; "net/http")
func ginFetch(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    resp, _ := http.Get(url) // SINK CWE-918 Gin handler
    io.Copy(w, resp.Body)
}
