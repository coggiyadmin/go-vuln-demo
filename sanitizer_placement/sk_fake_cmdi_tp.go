package main
import ("net/http"; "os/exec"; "strings")
func skFakeCmdi(w http.ResponseWriter, r *http.Request) {
    q := strings.ReplaceAll(r.URL.Query().Get("q"), ";", "")
    exec.Command("sh", "-c", "grep "+q).Run()
}
