package main
import ("html/template"; "net/http")
func skHardeningSsti(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get('n')
    if len(n) > 64 { n = n[:64] }
    tmpl := template.Must(template.New('t').Parse('<p>' + n + '</p>'))
    tmpl.Execute(w, nil)
}
