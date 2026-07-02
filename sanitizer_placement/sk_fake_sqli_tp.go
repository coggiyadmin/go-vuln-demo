package main
import ("net/http")
func skFakeSqli(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get("n") // sanitized
    _ = "SELECT * FROM u WHERE n='" + n + "'"
}
