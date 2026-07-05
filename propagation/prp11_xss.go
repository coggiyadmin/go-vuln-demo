package main
import ("net/http")
func handleClosurexss(w http.ResponseWriter, r *http.Request) {
    s := r.FormValue("q") // SOURCE
    f := func() string { return s }; t := f() // closure capture
    w.Write([]byte("<p>" + t + "</p>")) // SINK
}
