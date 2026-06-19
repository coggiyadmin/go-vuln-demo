// Combination #5 — OOP × XSS (CWE-79, Go). NO finding = FN (#78).
package main

import "net/http"

type page struct{ t string }

func (p page) render() string { return "<h1>" + p.t + "</h1>" } // CWE-79

func oopXssHandler(w http.ResponseWriter, r *http.Request) {
	p := page{t: r.FormValue("q")}
	w.Write([]byte(p.render()))
}
