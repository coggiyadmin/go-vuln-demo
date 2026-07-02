package main
import ("fmt"; "html"; "net/http")
func skFrameworkNativeFmtstr(w http.ResponseWriter, r *http.Request) {
    pat := r.URL.Query().Get("fmt")
    name := html.EscapeString(r.URL.Query().Get("name"))
    fmt.Fprintf(w, pat, name)
}
