package loginjxf

import "net/http"

func Handle(r *http.Request) { Run(r.FormValue("user")) } // SOURCE -> cross-file sink (CWE-117)
