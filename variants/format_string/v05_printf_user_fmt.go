package main
import ("fmt"; "net/http")
func printfFmt(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user") // SOURCE
    fmt.Fprintf(w, user, "ok") // SINK CWE-134 format string — user controls format
}
