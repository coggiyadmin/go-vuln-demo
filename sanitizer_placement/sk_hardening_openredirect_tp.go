package main
import ("net/http"; "strings")
func skHardeningOpenredirect(w http.ResponseWriter, r *http.Request) {
    nxt := r.URL.Query().Get("next")
    if !strings.HasPrefix(nxt, "http://") && !strings.HasPrefix(nxt, "https://") {
        nxt = "https://example.com/" + nxt
    }
    http.Redirect(w, r, nxt, http.StatusFound)
}
