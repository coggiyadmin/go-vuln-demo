package main
import ("net/http"; "os"; "strings")
func skFakePathtrav(w http.ResponseWriter, r *http.Request) {
    p := strings.TrimPrefix(r.URL.Query().Get("p"), "/")
    os.ReadFile("/data/" + p)
}
