package main
import ("html"; "net/http")
func skWrongContextSqli(w http.ResponseWriter, r *http.Request) {
    n := html.EscapeString(r.URL.Query().Get("n"))
    _ = "SELECT * FROM u WHERE n='" + n + "'"
}
