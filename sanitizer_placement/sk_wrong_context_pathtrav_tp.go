package main
import ("html"; "net/http"; "os")
func skWrongContextPathtrav(w http.ResponseWriter, r *http.Request) {
    p := html.EscapeString(r.URL.Query().Get("p"))
    os.ReadFile("/data/" + p)
}
