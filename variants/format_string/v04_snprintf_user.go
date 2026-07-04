package main
import ("fmt"; "net/http")
func snprintfFmt(w http.ResponseWriter, r *http.Request) {
    user := r.URL.Query().Get("user")
    msg := fmt.Sprintf("user=%s action=login", user) // SINK CWE-134 format string
    w.Write([]byte(msg))
}
