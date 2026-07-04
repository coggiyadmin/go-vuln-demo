package main
import ("net/http"; "net/url")
var allowed = map[string]bool{"app.example.com": true}
func safeGinRedirect(w http.ResponseWriter, r *http.Request) {
    nxt := r.URL.Query().Get("next")
    u, _ := url.Parse(nxt)
    if allowed[u.Hostname()] {
        http.Redirect(w, r, nxt, http.StatusFound)
        return
    }
    http.Redirect(w, r, "/home", http.StatusFound)
}
