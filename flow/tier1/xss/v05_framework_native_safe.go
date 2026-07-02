package main
import ("html"; "net/http")
func v05(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    w.Write([]byte("<p>" + html.EscapeString(q) + "</p>")) // SK-05 encode
}
