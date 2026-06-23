// Excessive Agency FN (OWASP LLM06) — confused deputy: model-chosen URL fetched with an
// ambient privileged token, no host allowlist (SSRF via tool). Expected: trust layer MISS.
package aisec

import (
	"io"
	"net/http"
	"os"
)

func FetchTool(url string) (string, error) {
	req, _ := http.NewRequest("GET", url, nil) // SINK (LLM06 confused-deputy)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_TOKEN"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return string(b), nil
}
