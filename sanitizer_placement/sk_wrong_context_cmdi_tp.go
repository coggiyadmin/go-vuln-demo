package main
import ("net/http"; "net/url"; "os/exec")
func skWrongContextCmdi(w http.ResponseWriter, r *http.Request) {
    q := url.QueryEscape(r.URL.Query().Get("q"))
    exec.Command("sh", "-c", "grep "+q).Run()
}
