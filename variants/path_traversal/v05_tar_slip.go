package main
import ("archive/tar"; "io"; "net/http"; "os")
func tarSlip(w http.ResponseWriter, r *http.Request) {
    tr := tar.NewReader(os.Stdin)
    tr.Next()
    f, _ := os.Create("/var/data/" + r.FormValue("name")) // SINK CWE-22 tar slip
    io.Copy(f, tr)
}
