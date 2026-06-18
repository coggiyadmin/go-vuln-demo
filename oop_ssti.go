// Combination #5 — OOP OBJECT FLOW × SSTI (CWE-1336, Go). Tainted struct field
// concatenated into a text/template source. NO finding = FALSE NEGATIVE.
package main

import (
	"net/http"
	"os"
	"text/template"
)

type greetTmpl struct{ name string }

func (g greetTmpl) renderDirect() {
	t, _ := template.New("t").Parse("<p>Hello " + g.name + "</p>") // template-injection sink (CWE-1336)
	t.Execute(os.Stdout, nil)
}

func sstiOOPHandler(w http.ResponseWriter, r *http.Request) {
	g := greetTmpl{name: r.FormValue("name")} // SOURCE -> struct field
	g.renderDirect()
}
