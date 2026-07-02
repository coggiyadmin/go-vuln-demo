// Phase-3 encode mirror — contextual encoding before sink
package main
import ("html/template"; "net/http")
func sstiSafe(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("t").Parse("<p>{{.}}</p>"))
    tmpl.Execute(w, r.FormValue("n"))
}
