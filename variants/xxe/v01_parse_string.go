package main
import ("encoding/xml"; "net/http")
func xxeParse(w http.ResponseWriter, r *http.Request) {
    var v interface{}
    xml.NewDecoder(r.Body).Decode(&v) // SINK CWE-611
}
