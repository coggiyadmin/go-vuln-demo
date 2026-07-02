package main
import ("encoding/xml"; "io"; "net/http"; "strings")
func companySanitize(b []byte) []byte {
    return []byte(strings.ReplaceAll(string(b), "ENTITY", ""))
}
func skCustomWrapperXxe(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    var v interface{}
    xml.NewDecoder(strings.NewReader(string(companySanitize(body)))).Decode(&v)
}
