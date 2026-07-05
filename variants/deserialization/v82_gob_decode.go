package main
import ("encoding/gob"; "net/http")
func decode(w http.ResponseWriter, r *http.Request) {
    var v map[string]interface{}
    gob.NewDecoder(r.Body).Decode(&v) // SINK CWE-502 gob decode of untrusted stream
}
