// SAFE mirror — oop_xss.go. Expect 0.
package main

import (
	"html"
	"net/http"
)

type safePage struct{ t string }

func (p safePage) render() string { return "<h1>" + html.EscapeString(p.t) + "</h1>" }

func safeOopXssHandler(w http.ResponseWriter, r *http.Request) {
	p := safePage{t: r.FormValue("q")}
	w.Write([]byte(p.render()))
}
