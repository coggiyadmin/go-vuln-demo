package main
import "net/http"
func redir(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url") // SOURCE
    http.Redirect(w, r, url, http.StatusFound) // SINK CWE-601 go http.Redirect
}
