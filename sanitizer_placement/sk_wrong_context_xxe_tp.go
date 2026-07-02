package main
import ("encoding/xml"; "html"; "net/http")
func skWrongContextXxe(w http.ResponseWriter, r *http.Request) {
    _ = html.EscapeString(r.URL.Query().Get("n"))
    var v interface{}
    xml.NewDecoder(r.Body).Decode(&v)
}
