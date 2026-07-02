package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, '\r', "") }
func skCustomWrapperCrlf(w http.ResponseWriter, r *http.Request) {
    nxt := companySanitize(r.URL.Query().Get('next'))
    w.Header().Set('Location', '/home?u='+nxt)
}
