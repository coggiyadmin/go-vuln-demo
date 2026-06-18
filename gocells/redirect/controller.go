package redirxf

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("next") // SOURCE
	Run(w, r, url)             // cross-file sink (CWE-601)
}
