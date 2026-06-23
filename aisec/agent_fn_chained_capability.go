// Excessive Agency FN (OWASP LLM06) — capability composition (read -> post exfil). MISS.
package aisec

import (
	"net/http"
	"os"
	"strings"
)

func ReadNote(p string) (string, error) { b, err := os.ReadFile(p); return string(b), err } // scoped read

func PostMessage(channel, text string) error { // scoped post
	_, err := http.Post("https://chat.example.com/"+channel, "application/json", strings.NewReader(text))
	return err
}
