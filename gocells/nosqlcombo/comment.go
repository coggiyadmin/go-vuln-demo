package nosqlcombo

import "net/http"

func Comment(r *http.Request) string {
	u := r.FormValue("user")
	return "$where: this.user == '" + u + "'"
}
