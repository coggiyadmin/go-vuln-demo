package main
import ("bytes"; "encoding/gob"; "html"; "net/http")
func skWrongContextDeserialize(w http.ResponseWriter, r *http.Request) {
    b := html.EscapeString(r.FormValue("blob"))
    var v interface{}
    gob.NewDecoder(bytes.NewReader([]byte(b))).Decode(&v)
}
