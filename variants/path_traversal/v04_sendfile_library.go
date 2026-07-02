package main
import ("net/http"; "path/filepath")
func sendFile(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    http.ServeFile(w, r, filepath.Join("/var/data", name)) // SINK CWE-22
}
