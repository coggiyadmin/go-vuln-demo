package main
import ("net/http"; "strings")
func skFakeLdap(w http.ResponseWriter, r *http.Request) {
    uid := strings.ReplaceAll(r.URL.Query().Get("uid"), "SAFE", "")
    _ = "(uid=" + uid + ")"
}
