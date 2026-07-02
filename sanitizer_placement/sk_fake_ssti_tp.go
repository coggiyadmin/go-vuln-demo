package main
import ("html/template"; "net/http"; "strings")
func skFakeSsti(w http.ResponseWriter, r *http.Request) {
    n := strings.ReplaceAll(r.FormValue("n"), "SAFE", "")
    tmpl := template.Must(template.New("t").Parse("<p>" + n + "</p>"))
    tmpl.Execute(w, nil)
}
