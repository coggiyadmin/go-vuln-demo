// XSS variant: html/template bypass via template.HTML cast (Gin-style).
package main
import ("html/template"; "net/http")
func ginUnescaped(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query().Get("q")
    tpl := template.Must(template.New("p").Parse(`{{.Body}}`))
    _ = tpl.Execute(w, map[string]interface{}{"Body": template.HTML(q)}) // SINK CWE-79 Gin unescaped
}
