package main
import ("io"; "net/http")
func fastapiProxy(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    resp, _ := http.Get(url) // SINK CWE-918
    io.Copy(w, resp.Body)
}
