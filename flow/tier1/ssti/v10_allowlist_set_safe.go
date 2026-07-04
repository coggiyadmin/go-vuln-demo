// v10 allowlist-set
package main
import "html/template"
func render(name string) (string, error) {
	t, err := template.New("p").Parse("<p>{{.}}</p>"); if err != nil { return "", err }
	var b []byte; t.Execute(&b, name); return string(b), nil
}
