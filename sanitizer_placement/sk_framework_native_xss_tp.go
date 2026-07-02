package main
import ("html/template"; "net/http")
func skFrameworkNativeXss(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    tpl := template.Must(template.New("x").Parse("<p>{{.}}</p>"))
    tpl.Execute(w, template.HTML(q))
}
