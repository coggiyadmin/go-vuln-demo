package main
import ("html"; "net/http")
func skFrameworkNativeCodeinj(w http.ResponseWriter, r *http.Request) {
    expr := html.EscapeString(r.URL.Query().Get('expr'))
    _ = 'eval(' + expr + ')'
}
