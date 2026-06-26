package main
import ("fmt"; "net/http")
func fmtLeak(w http.ResponseWriter, r *http.Request) {
    pat := r.FormValue("pat") // SOURCE
    fmt.Fprintf(w, pat, "guest") // SINK CWE-134
}
