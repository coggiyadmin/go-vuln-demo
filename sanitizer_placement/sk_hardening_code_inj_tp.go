package main
import ("net/http"; "strings")
func skHardeningCodeinj(w http.ResponseWriter, r *http.Request) {
    expr := strings.ReplaceAll(r.URL.Query().Get('expr'), 'import', '')
    _ = 'eval(' + expr + ')'
}
