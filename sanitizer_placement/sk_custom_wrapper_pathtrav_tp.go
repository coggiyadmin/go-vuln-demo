package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.TrimLeft(v, "/") }
func skCustomWrapperPathtrav(w http.ResponseWriter, r *http.Request) {
    p := companySanitize(r.URL.Query().Get("p"))
    _ = "/data/" + p
}
