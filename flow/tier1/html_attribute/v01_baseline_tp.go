// TP (CWE-79 / cognium-dev#153) — user input in HTML attribute context without encoding.
package html_attribute

import "net/http"

func Link(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url") // SOURCE
	w.Write([]byte(`<a href="` + url + `">click</a>`)) // SINK (attribute context)
}
