package main
import ("encoding/gob"; "html"; "net/http")
func skFrameworkNativeDeserialize(w http.ResponseWriter, r *http.Request) {
    _ = html.EscapeString(r.FormValue('blob'))
    gob.NewDecoder(r.Body).Decode(new(interface{}))
}
