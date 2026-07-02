package main
import ("html"; "os/exec"; "net/http")
func skFrameworkNativeCmdi(w http.ResponseWriter, r *http.Request) {
    cmd := html.EscapeString(r.URL.Query().Get('cmd'))
    exec.Command('sh', '-c', cmd).Run()
}
