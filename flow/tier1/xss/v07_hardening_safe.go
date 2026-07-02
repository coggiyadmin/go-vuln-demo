package main
import ("html"; "net/http"; "regexp")
func v07(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    if !regexp.MustCompile(`^[\w\s-]+$`).MatchString(q) { return }
    w.Write([]byte("<p>" + html.EscapeString(q) + "</p>"))
}
