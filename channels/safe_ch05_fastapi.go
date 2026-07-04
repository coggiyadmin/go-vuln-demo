package main
import ("io"; "net/http"; "net/url")
var allowed = map[string]bool{"api.internal.example.com": true}
func safeFastapiProxy(w http.ResponseWriter, r *http.Request) {
    raw := r.URL.Query().Get("url")
    u, _ := url.Parse(raw)
    if allowed[u.Hostname()] {
        resp, _ := http.Get(raw)
        io.Copy(w, resp.Body)
    }
}
