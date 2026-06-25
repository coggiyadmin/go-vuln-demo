// TP — wiki macro evaluates user-supplied template source (CWE-94 intake pattern).
// CVE-2025-24893 class: server-side template/code injection in generic wiki platform.
package main

import (
	"net/http"
	"text/template"
)

func wikiRender(w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("content")
	t, _ := template.New("wiki").Parse(src) // SINK CWE-94 — user controls template program
	_ = t.Execute(w, nil)
}
