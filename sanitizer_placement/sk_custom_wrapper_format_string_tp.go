package main
import ("fmt"; "net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, "%", "") }
func skCustomWrapperFmtstr(w http.ResponseWriter, r *http.Request) {
    pat := companySanitize(r.URL.Query().Get("fmt"))
    fmt.Fprintf(w, pat, "guest")
}
