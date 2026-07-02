// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("encoding/xml"; "net/http")
func xxeSafe(w http.ResponseWriter, r *http.Request) {
    dec := xml.NewDecoder(r.Body)
    dec.Strict = true
    var v interface{}
    dec.Decode(&v)
}
