package main
import ("net/http"; "os/exec"; "strings")
func skHardeningCmdi(w http.ResponseWriter, r *http.Request) {
    cmd := strings.ReplaceAll(r.URL.Query().Get('cmd'), ';', '')
    exec.Command('sh', '-c', cmd).Run()
}
