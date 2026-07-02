package main
import ("net/http"; "strings")
func skHardeningSsrf(w http.ResponseWriter, r *http.Request) {
    u := r.URL.Query().Get("url")
    if strings.Contains(u, "127.0.0.1") || strings.Contains(u, "169.254.169.254") {
        http.Error(w, "blocked", http.StatusForbidden)
        return
    }
    http.Get(u)
}
