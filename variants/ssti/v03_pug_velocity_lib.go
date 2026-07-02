package main
import ("html/template"; "net/http")
func pugStyle(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.New("t").Parse("Hello " + r.FormValue("name")))
    tmpl.Execute(w, nil) // SINK CWE-94 text/template concat
}
