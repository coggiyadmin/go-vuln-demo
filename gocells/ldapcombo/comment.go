package ldapcombo

import "net/http"

func Comment(r *http.Request) string {
	u := r.FormValue("user")
	ex := "(uid=" + u + ")"
	return ex
}
