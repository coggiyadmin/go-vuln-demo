package pathxf

import "net/http"

func Handle(r *http.Request) { name := r.FormValue("name"); Run(name) } // SOURCE -> cross-file sink (CWE-22)
