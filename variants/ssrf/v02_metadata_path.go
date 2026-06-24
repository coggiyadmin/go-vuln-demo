package main
import ("net/http"; "io")
func ssrfMeta(w http.ResponseWriter, r *http.Request) {
    p := r.FormValue("path")
    resp, _ := http.Get("http://169.254.169.254/" + p) // SINK CWE-918
    io.Copy(w, resp.Body)
}
