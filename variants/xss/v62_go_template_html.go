package main
import ("html/template"; "net/http")
func render(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name") // SOURCE
    t := template.Must(template.New("p").Parse("<h1>{{.}}</h1>"))
    t.Execute(w, template.HTML(name)) // SINK CWE-79 template.HTML bypasses auto-escape
}
