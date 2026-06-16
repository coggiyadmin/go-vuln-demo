package xf
import "net/http"
func Handle(r *http.Request){ h:=r.FormValue("h"); Run(h) } // SOURCE -> cross-file sink
