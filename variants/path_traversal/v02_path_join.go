package main
import ("net/http"; "path/filepath"; "os")
func pathJoin(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    os.ReadFile(filepath.Join("/var/data", name)) // SINK CWE-22
}
