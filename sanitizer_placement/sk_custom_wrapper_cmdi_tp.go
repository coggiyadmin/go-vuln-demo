package main
import ("net/http"; "os/exec"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, ";", "") }
func skCustomWrapperCmdi(w http.ResponseWriter, r *http.Request) {
    q := companySanitize(r.URL.Query().Get("q"))
    exec.Command("sh", "-c", "grep "+q).Run()
}
