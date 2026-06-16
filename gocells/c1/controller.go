package c1
import "net/http"
func Handle(r *http.Request){ name := r.FormValue("h"); Run(name) } // SOURCE -> cross-file sink
