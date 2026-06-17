// FN probe — CWE-79 reflected XSS (Go category probe).
package main

import (
	"fmt"
	"net/http"
)

func xssGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name)
}
