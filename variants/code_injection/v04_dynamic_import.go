package main
import ("plugin"; "net/http")
func codePlugin(w http.ResponseWriter, r *http.Request) {
    path := r.FormValue("path")
    plugin.Open(path) // SINK CWE-94 dynamic load
}
