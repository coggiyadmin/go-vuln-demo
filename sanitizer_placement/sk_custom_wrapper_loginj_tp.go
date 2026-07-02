package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, '\n', "") }
func skCustomWrapperLoginj(w http.ResponseWriter, r *http.Request) {
    user := companySanitize(r.URL.Query().Get('user'))
    _ = 'login user=' + user
}
