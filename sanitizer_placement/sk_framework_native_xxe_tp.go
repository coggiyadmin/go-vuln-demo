package main
import ("encoding/xml"; "html"; "net/http")
func skFrameworkNativeXxe(w http.ResponseWriter, r *http.Request) {
    x := html.EscapeString(r.FormValue('xml'))
    var v interface{}
    _ = xml.Unmarshal([]byte(x), &v)
}
