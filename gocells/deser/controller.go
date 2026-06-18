package deserxf

import "net/http"

func Handle(r *http.Request) { Run(r.FormValue("s")) } // SOURCE -> cross-file sink (CWE-502)
