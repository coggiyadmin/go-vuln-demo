package main
import ("encoding/xml"; "net/http")
func libxmlStyle(w http.ResponseWriter, r *http.Request) {
    var v interface{}
    xml.Unmarshal([]byte(r.FormValue("xml")), &v) // SINK CWE-611 encoding/xml
}
