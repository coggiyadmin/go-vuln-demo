package main
import ("net/http"; "github.com/valyala/fasthttp")
func fetch(w http.ResponseWriter, r *http.Request) {
    url := r.FormValue("url") // SOURCE
    _, _, _ = fasthttp.Get(nil, url) // SINK CWE-918 fasthttp client
}
