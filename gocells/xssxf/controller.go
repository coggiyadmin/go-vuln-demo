package xssxf

import "net/http"

func Handle(r *http.Request) string { q := r.FormValue("q"); return Run(q) } // SOURCE -> cross-file (CWE-79)
