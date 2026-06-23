// Improper Output Handling (OWASP LLM05) — model output written to an HTTP response as raw HTML. TP.
package aisec

import "net/http"

func RenderAnswer(w http.ResponseWriter, answer string) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<div>" + answer + "</div>")) // SINK (LLM05->XSS)
}
