package main
import ("net/http"; "regexp"; "strings")
func skFakeRedos(w http.ResponseWriter, r *http.Request) {
    pattern := strings.ReplaceAll(r.URL.Query().Get("p"), "SAFE", "")
    re, _ := regexp.Compile(pattern)
    re.MatchString(r.URL.Query().Get("t"))
}
