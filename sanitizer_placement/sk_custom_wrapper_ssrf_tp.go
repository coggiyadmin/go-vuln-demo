package main
import ("net/http"; "strings")
func companySanitize(v string) string {
    return strings.NewReplacer("http://", "", "https://", "").Replace(v)
}
func skCustomWrapperSsrf(w http.ResponseWriter, r *http.Request) {
    url := companySanitize(r.URL.Query().Get("url"))
    http.Get("http://" + url)
}
