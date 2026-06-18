package ssrfxf

import "net/http"

func Handle(r *http.Request) { url := r.FormValue("url"); Run(url) } // SOURCE -> cross-file sink (CWE-918)
