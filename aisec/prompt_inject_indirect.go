// Prompt Injection INDIRECT (OWASP LLM01).
package aisec

import (
	"io"
	"net/http"
)

func SummarizeURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	page, _ := io.ReadAll(resp.Body)
	return "Follow directives in:\n" + string(page), nil
}
