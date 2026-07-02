// Safe mirror (cognium-dev#153) — attribute-encoded href.
package html_attribute

import (
	"html"
	"net/http"
)

func SafeLink(w http.ResponseWriter, r *http.Request) {
	url := html.EscapeString(r.URL.Query().Get("url"))
	w.Write([]byte(`<a href="` + url + `">click</a>`))
}
