package main
import ("io"; "net/http")
func nativeClient(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    resp, _ := http.DefaultClient.Get(url) // SINK CWE-918 net/http DefaultClient
    io.Copy(w, resp.Body)
}
