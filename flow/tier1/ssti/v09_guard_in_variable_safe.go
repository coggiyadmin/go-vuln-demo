// v09 guard-in-variable
package main
import "html/template"
var allowed = map[string]bool{"Hello {{.Name}}": true}
func render(tmpl string, name string) string { if allowed[tmpl] { t, _ := template.New("t").Parse(tmpl); var b []byte; t.Execute(&b, map[string]string{"Name": name}); return string(b) }; return "" }
