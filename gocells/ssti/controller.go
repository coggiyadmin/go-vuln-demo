package sstixf

import "net/http"

func Handle(r *http.Request) { Run(r.FormValue("name")) } // SOURCE -> cross-file sink (CWE-1336)
