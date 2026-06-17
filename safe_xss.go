// SAFE mirror — xss_probe.go; HTML-encode before write.
package main

import (
	"fmt"
	"html"
	"net/http"
)

func safeXssGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	escaped := html.EscapeString(name)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", escaped)
}
