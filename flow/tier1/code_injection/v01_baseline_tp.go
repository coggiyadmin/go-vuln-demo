package main
import ("github.com/robertkrimen/otto"; "net/http")
func codeInj(w http.ResponseWriter, r *http.Request) {
    x := r.URL.Query().Get("x") // SOURCE
    vm := otto.New(); vm.Run(x) // SINK CWE-94
}
