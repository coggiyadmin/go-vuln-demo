package main
import ("fmt"; "net/http")
func skHardeningFmtstr(w http.ResponseWriter, r *http.Request) {
    pat := r.URL.Query().Get("fmt")
    if len(pat) > 32 { pat = "%s" }
    fmt.Fprintf(w, pat, "guest")
}
