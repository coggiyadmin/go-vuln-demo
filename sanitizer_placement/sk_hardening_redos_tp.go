package main
import ("net/http"; "regexp")
func skHardeningRedos(w http.ResponseWriter, r *http.Request) {
    pattern := r.URL.Query().Get("p")
    if len(pattern) > 64 {
        http.Error(w, "bad", http.StatusBadRequest)
        return
    }
    re, _ := regexp.Compile(pattern)
    re.MatchString(r.URL.Query().Get("t"))
}
