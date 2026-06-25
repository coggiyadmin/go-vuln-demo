// Safe mirror — fixed template; user data passed only as Execute context value.
package main

import (
	"html/template"
	"net/http"
)

func wikiRenderSafe(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("content")
	t := template.Must(template.New("wiki").Parse("<div class=\"wiki-body\">{{.}}</div>"))
	_ = t.Execute(w, body)
}
