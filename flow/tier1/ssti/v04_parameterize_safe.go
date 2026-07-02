// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("html/template"; "net/http")
func sstiSafe(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("t").Parse("<p>{{.}}</p>"))
    tmpl.Execute(w, r.FormValue("n"))
}
