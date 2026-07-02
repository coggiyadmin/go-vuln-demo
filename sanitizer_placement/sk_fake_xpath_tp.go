package main
import ("net/http"; "strings")
func skFakeXpath(w http.ResponseWriter, r *http.Request) {
    name := strings.ReplaceAll(r.URL.Query().Get("name"), "SAFE", "")
    _ = "//user[name='" + name + "']"
}
