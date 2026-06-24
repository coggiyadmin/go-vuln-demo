package main
import ("io"; "net/http"; "os")
func pathRead(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    f, _ := os.Open("/var/data/" + name) // SINK CWE-22
    io.Copy(w, f)
}
