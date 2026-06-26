package main
import ("net/http"; "github.com/antchfx/htmlquery")
func xpathConcat(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name") // SOURCE
    htmlquery.FindOne(nil, "//user[name='"+name+"']") // SINK CWE-643
}
