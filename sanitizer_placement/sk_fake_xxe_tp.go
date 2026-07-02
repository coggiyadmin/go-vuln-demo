package main
import ("encoding/xml"; "net/http"; "strings")
func skFakeXxe(w http.ResponseWriter, r *http.Request) {
    _ = strings.ReplaceAll(r.URL.Query().Get("tag"), "SAFE", "")
    var v interface{}
    xml.NewDecoder(r.Body).Decode(&v)
}
