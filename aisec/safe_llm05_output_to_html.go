// SAFE mirror (OWASP LLM05) — model output HTML-escaped before writing.
package aisec

import (
	"html"
	"net/http"
)

func RenderAnswer(w http.ResponseWriter, answer string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<div>" + html.EscapeString(answer) + "</div>"))
}
