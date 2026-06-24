package main
import ("html/template"; "net/http")
func sstiTmpl(w http.ResponseWriter, r *http.Request) {
    n := r.FormValue("n")
    tmpl := template.Must(template.New("t").Parse("<p>" + n + "</p>")) // SINK CWE-1336
    tmpl.Execute(w, nil)
}
