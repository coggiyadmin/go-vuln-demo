package main
import ("net/http")
func openRedir(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, r.FormValue("next"), http.StatusFound) // SINK CWE-601
}
