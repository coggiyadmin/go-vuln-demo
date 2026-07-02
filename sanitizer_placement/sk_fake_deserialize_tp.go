package main
import ("bytes"; "encoding/gob"; "net/http"; "strings")
func skFakeDeserialize(w http.ResponseWriter, r *http.Request) {
    b := strings.ReplaceAll(r.FormValue("blob"), "SAFE", "")
    var v interface{}
    gob.NewDecoder(bytes.NewReader([]byte(b))).Decode(&v)
}
