// Z6 N2 benign mirror — html/template auto-escape (Gin-style).
package libapi

import ("html/template"; "net/http")

func RenderSafe(w http.ResponseWriter, body string) error {
	tpl := template.Must(template.New("p").Parse("<p>{{.}}</p>"))
	return tpl.Execute(w, body)
}
