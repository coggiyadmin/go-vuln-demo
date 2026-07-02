package main
import ("encoding/gob"; "net/http")
func skHardeningDeserialize(w http.ResponseWriter, r *http.Request) {
    if r.ContentLength < 0 || r.ContentLength > 1<<20 {
        http.Error(w, "bad", http.StatusBadRequest)
        return
    }
    gob.NewDecoder(r.Body).Decode(new(interface{}))
}
