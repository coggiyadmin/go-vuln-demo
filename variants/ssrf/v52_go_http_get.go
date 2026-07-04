package main
import ("io"; "net/http")
func fetch(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url") // SOURCE
    resp, _ := http.Get(url) // SINK CWE-918 http.Get
    io.Copy(w, resp.Body)
}
