package main
import "net/http"
func crlfLoc(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Location", r.FormValue("url")) // SINK CWE-93
}
