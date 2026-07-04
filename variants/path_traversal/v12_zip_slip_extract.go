package main
import ("archive/zip"; "io"; "net/http"; "os")
func zipSlip(w http.ResponseWriter, r *http.Request) {
    zr, _ := zip.NewReader(r.Body, r.ContentLength)
    f, _ := os.Create("/var/www/" + r.FormValue("name")) // SINK CWE-22 zip slip
    io.Copy(f, zr.File[0].Open())
}
