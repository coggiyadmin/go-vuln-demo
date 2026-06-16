package main
import ("net/http"; htmltmpl "html/template"; texttmpl "text/template")
func unsafeTmpl(w http.ResponseWriter, r *http.Request){
  name := r.FormValue("name")
  t := texttmpl.Must(texttmpl.New("p").Parse("<p>{{.}}</p>"))
  t.Execute(w, name)   // text/template: NO HTML escaping -> CWE-79 XSS
}
func safeTmpl(w http.ResponseWriter, r *http.Request){
  name := r.FormValue("name")
  t := htmltmpl.Must(htmltmpl.New("p").Parse("<p>{{.}}</p>"))
  t.Execute(w, name)   // html/template: auto-escapes -> SAFE
}
