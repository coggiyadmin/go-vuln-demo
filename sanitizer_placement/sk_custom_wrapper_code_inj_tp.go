package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, ";", "") }
func skCustomWrapperCodeinj(w http.ResponseWriter, r *http.Request) {
    code := companySanitize(r.URL.Query().Get("code"))
    _ = code
}
