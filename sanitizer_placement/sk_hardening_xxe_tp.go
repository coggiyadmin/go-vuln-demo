package main
import ("encoding/xml"; "net/http")
func skHardeningXxe(w http.ResponseWriter, r *http.Request) {
    if r.ContentLength > 65536 {
        http.Error(w, "too large", http.StatusRequestEntityTooLarge)
        return
    }
    var v interface{}
    xml.NewDecoder(r.Body).Decode(&v)
}
