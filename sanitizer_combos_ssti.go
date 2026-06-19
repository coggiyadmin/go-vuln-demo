// Combinations #6/#7/#8 — SANITIZER × SSTI (CWE-1336, Go).
package main

import (
	"html/template"
	"net/http"
	"strings"
)

func scStripHTMLSsti(s string) string { return strings.NewReplacer("<", "_").Replace(s) }
func scFakeSsti(s string) string      { return s }

func scWrongSsti(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("t").Parse("<p>" + scStripHTMLSsti(r.FormValue("name")) + "</p>"))
	t.Execute(w, nil) // CWE-1336
}

func scFakeSsti(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("t").Parse("<p>" + scFakeSsti(r.FormValue("name")) + "</p>"))
	t.Execute(w, nil) // CWE-1336
}

func scWrappedSsti(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("t").Parse("<p>{{.Name}}</p>"))
	t.Execute(w, map[string]string{"Name": r.FormValue("name")}) // expect 0
}
