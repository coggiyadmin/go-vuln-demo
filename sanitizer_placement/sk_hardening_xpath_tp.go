package main
import ("net/http")
func skHardeningXpath(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if len(name) > 256 {
        http.Error(w, "too long", http.StatusRequestEntityTooLarge)
        return
    }
    _ = "//user[name='" + name + "']"
}
