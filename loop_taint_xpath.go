// Combination #3 — LOOP-CARRIED TAINT × XPATH (CWE-643, Go).
package main

import (
	"net/http"
	"github.com/antchfx/htmlquery"
)

func ltXpathList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["uid"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	uid := items[0]
	htmlquery.FindOne(nil, "//user[name='"+uid+"']") // CWE-643
}
