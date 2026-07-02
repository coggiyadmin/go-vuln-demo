package main
import ("html"; "net/http")
func v06(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    w.Write([]byte(renderSafe(q)))
}
func renderSafe(q string) string { return "<p>" + html.EscapeString(q) + "</p>" }
