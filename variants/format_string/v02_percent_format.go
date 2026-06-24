package main
import ("fmt"; "net/http")
func sprintf(w http.ResponseWriter, r *http.Request) {
    msg := r.FormValue("msg")
    fmt.Sprintf("%s", msg) // SINK CWE-134
}
