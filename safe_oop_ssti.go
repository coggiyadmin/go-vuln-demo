// SAFE mirror — oop_ssti.go; constant template, user value passed as data. Expect 0 findings.
package main

import (
	"net/http"
	"text/template"
)

type safeGreetTmpl struct{ name string }

func (g safeGreetTmpl) renderDirect(w http.ResponseWriter) {
	t := template.Must(template.New("t").Parse("<p>Hello {{.}}</p>")) // constant template
	t.Execute(w, g.name)                                              // value as data
}

func safeSstiOOPHandler(w http.ResponseWriter, r *http.Request) {
	g := safeGreetTmpl{name: r.FormValue("name")}
	g.renderDirect(w)
}
