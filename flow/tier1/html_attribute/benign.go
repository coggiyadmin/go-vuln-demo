// Benign — static href.
package html_attribute

import "net/http"

func Benign(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(`<a href="/home">home</a>`))
}
