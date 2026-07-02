package main
import ("net/http"; "strings")
func companySanitize(v string) string {
    return strings.ReplaceAll(strings.ReplaceAll(v, "<", ""), ">", "")
}
func skCustomWrapperXss(w http.ResponseWriter, r *http.Request) {
    q := companySanitize(r.URL.Query().Get("q"))
    w.Write([]byte("<p>" + q + "</p>"))
}
