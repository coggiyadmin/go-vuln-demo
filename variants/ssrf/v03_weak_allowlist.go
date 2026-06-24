package main
import ("net/http"; "io"; "strings")
func ssrfWeak(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url")
    if strings.Contains(url, "trusted") {
        resp, _ := http.Get(url) // SINK CWE-918
        io.Copy(w, resp.Body)
    }
}
