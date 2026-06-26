package main
import "net/http"
func xss(w http.ResponseWriter, r *http.Request) {
    v := r.URL.Query().Get("q") // SOURCE
    w.Write([]byte("<p>" + v + "</p>")) // SINK CWE-79
}
