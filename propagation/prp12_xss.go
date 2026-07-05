package main
import ("net/http")
var gStore string
func handleGlobalVarxss(w http.ResponseWriter, r *http.Request) {
    s := r.FormValue("q") // SOURCE
    gStore = s; t := gStore // package global
    w.Write([]byte("<p>" + t + "</p>")) // SINK
}
