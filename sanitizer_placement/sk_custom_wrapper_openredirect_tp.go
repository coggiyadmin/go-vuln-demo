package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(strings.ToLower(v), "javascript:", "") }
func skCustomWrapperOpenredirect(w http.ResponseWriter, r *http.Request) {
    nxt := companySanitize(r.URL.Query().Get("next"))
    http.Redirect(w, r, nxt, http.StatusFound)
}
