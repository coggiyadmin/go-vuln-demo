package main
import "net/http"
func ssrfBlind(w http.ResponseWriter, r *http.Request) {
    http.Get(r.FormValue("cb")) // SINK CWE-918
}
