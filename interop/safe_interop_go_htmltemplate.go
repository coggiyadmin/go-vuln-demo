// IL-1 polyglot — SAFE mirror of interop_go_htmltemplate.go.
// The untrusted value is passed as a plain string, so html/template applies its
// contextual auto-escaping (no template.HTML cast). The scanner MUST produce ZERO
// security findings.
package interop

import (
	"html/template"
	"net/http"
)

func SafeHandle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // SOURCE
	t := template.Must(template.New("p").Parse("<h1>{{.}}</h1>"))
	// Safe: passed as a string — html/template auto-escapes it; `<script>` becomes
	// &lt;script&gt;. No template.HTML cast, so escaping is not disabled.
	_ = t.Execute(w, name)
}
