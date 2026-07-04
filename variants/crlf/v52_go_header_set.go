package main
import "net/http"
func setHeader(w http.ResponseWriter, r *http.Request) {
    lang := r.FormValue("lang") // SOURCE
    w.Header().Set("Content-Language", lang) // SINK CWE-113 go header set
}
