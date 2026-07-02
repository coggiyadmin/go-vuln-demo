package main
import ("net/http"; "path/filepath")
func skHardeningPathtrav(w http.ResponseWriter, r *http.Request) {
    p := filepath.Base(r.URL.Query().Get("p"))
    _ = "/data/" + p
}
