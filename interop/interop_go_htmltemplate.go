// IL-1 polyglot — Go → html/template (CWE-79).
// Host (Go) embeds a guest language (HTML template) and bypasses the engine's
// contextual auto-escaping by casting untrusted input to template.HTML, so the
// value is emitted as raw markup. Expected today: FN (cast-bypass not modeled).
package interop

import (
	"html/template"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // SOURCE (HTTP param)
	t := template.Must(template.New("p").Parse("<h1>{{.}}</h1>"))
	// SINK (CWE-79): template.HTML cast tells html/template the value is already
	// safe markup, disabling auto-escaping — XSS for attacker-controlled `name`.
	_ = t.Execute(w, template.HTML(name))
}
