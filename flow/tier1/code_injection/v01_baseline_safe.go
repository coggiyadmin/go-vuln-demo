package main
import ("net/http")
func codeInjSafe(w http.ResponseWriter, r *http.Request) {
    x := r.URL.Query().Get("x")
    if x != "daily" && x != "weekly" { http.Error(w, "forbidden", 403); return }
}
