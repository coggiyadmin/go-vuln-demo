package main
import ("fmt"; "net/http")
func fmtSafe(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello %s", r.FormValue("name"))
}
