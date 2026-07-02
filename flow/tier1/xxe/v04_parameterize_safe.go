// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("encoding/xml"; "net/http")
func xxeSafe(w http.ResponseWriter, r *http.Request) {
    dec := xml.NewDecoder(r.Body)
    dec.Strict = true
    var v interface{}
    dec.Decode(&v)
}
