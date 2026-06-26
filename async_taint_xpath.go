// Combination #4 — ASYNC taint × XPATH (CWE-643, Go).
package main

import (
	"net/http"
	"github.com/antchfx/htmlquery"
)

func atXpathHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	ch := make(chan string, 1)
	go func() { ch <- uid }()
	v := <-ch
	htmlquery.FindOne(nil, "//user[name='"+uid+"']") // CWE-643
}
