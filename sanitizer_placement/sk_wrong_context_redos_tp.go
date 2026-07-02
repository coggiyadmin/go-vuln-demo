package main
import ("net/http"; "regexp"; "strings")
func skWrongContextRedos(w http.ResponseWriter, r *http.Request) {
    pattern := strings.ReplaceAll(r.URL.Query().Get("p"), "'", "''")
    re, _ := regexp.Compile(pattern)
    re.MatchString(r.URL.Query().Get("t"))
}
