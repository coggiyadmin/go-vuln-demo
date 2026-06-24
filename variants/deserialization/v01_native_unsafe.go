package main
import ("encoding/gob"; "bytes"; "net/http")
func deserGob(w http.ResponseWriter, r *http.Request) {
    var v interface{}
    gob.NewDecoder(bytes.NewReader([]byte(r.FormValue("blob")))).Decode(&v) // SINK CWE-502
}
