package xpathcombo

import "net/http"

func Comment(r *http.Request) string {
	n := r.FormValue("name")
	return "//user[name='" + n + "']"
}
