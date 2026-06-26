package main
import ("io/ioutil"; "net/http")
func pathTrav(w http.ResponseWriter, r *http.Request) {
    p := r.URL.Query().Get("p") // SOURCE
    ioutil.ReadFile("/data/" + p) // SINK CWE-22
}
