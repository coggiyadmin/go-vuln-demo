package main
import ("html"; "net/http")
func skWrongContextXpath(w http.ResponseWriter, r *http.Request) {
    name := html.EscapeString(r.URL.Query().Get("name"))
    _ = "//user[name='" + name + "']"
}
