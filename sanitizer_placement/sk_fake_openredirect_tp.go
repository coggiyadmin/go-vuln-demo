package main
import ("net/http"; "strings")
func skFakeOpenredirect(w http.ResponseWriter, r *http.Request) {
    nxt := strings.ReplaceAll(r.URL.Query().Get("next"), "SAFE", "")
    http.Redirect(w, r, nxt, http.StatusFound)
}
