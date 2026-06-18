package redirxf

import "net/http"

func Run(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusFound) // SINK CWE-601
}
