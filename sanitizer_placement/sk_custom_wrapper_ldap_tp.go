package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(strings.ReplaceAll(v, "(", ""), ")", "") }
func skCustomWrapperLdap(w http.ResponseWriter, r *http.Request) {
    uid := companySanitize(r.URL.Query().Get("uid"))
    _ = "(uid=" + uid + ")"
}
