package main
import ("html/template"; "net/http"; "strings")
func companySanitize(v string) string {
    return strings.ReplaceAll(strings.ReplaceAll(v, "{{", ""), "}}", "")
}
func skCustomWrapperSsti(w http.ResponseWriter, r *http.Request) {
    n := companySanitize(r.FormValue("n"))
    tmpl := template.Must(template.New("t").Parse("<p>" + n + "</p>"))
    tmpl.Execute(w, nil)
}
