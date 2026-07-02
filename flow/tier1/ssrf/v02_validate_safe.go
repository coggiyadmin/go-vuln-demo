// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("net/http"; "net/url")
func ssrfSafe(w http.ResponseWriter, r *http.Request) {
    u, _ := url.Parse(r.FormValue("url"))
    if u.Hostname() != "api.internal.example.com" { http.Error(w, "forbidden", 403); return }
    http.Get(u.String())
}
