package main
import ("net/http"; "regexp")
func redosMatch(w http.ResponseWriter, r *http.Request) {
    pattern := r.URL.Query().Get("p") // SOURCE
    re, _ := regexp.Compile(pattern)
    re.MatchString(r.URL.Query().Get("t")) // SINK CWE-1333
}
