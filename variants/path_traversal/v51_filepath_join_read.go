package main
import ("net/http"; "os"; "path/filepath")
func readFile(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name") // SOURCE
    data, _ := os.ReadFile(filepath.Join("/var/data", name)) // SINK CWE-22 filepath.Join no clean
    w.Write(data)
}
