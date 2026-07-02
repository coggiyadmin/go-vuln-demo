// Phase-3 parameterize mirror — bound params / prepared API
package main
import "net/http"
var allowed = map[string]bool{"/dashboard": true, "/profile": true}
func openRedirSafe(w http.ResponseWriter, r *http.Request) {
    nxt := r.FormValue("next")
    if !allowed[nxt] { http.Error(w, "forbidden", 403); return }
    http.Redirect(w, r, nxt, http.StatusFound)
}
