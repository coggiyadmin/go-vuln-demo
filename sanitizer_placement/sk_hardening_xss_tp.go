package main
import ("net/http"; "strings")
func skHardeningXss(w http.ResponseWriter, r *http.Request) {
    q := strings.ReplaceAll(r.URL.Query().Get('q'), '<script', '')
    w.Write([]byte('<p>' + q + '</p>'))
}
