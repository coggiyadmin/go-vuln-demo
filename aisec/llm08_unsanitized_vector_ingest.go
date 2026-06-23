// Vector/Embedding Weakness (OWASP LLM08) — untrusted doc embedded unsanitized. TP.
package aisec

import (
	"io"
	"net/http"
)

var vectorIndex []string

func Ingest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	vectorIndex = append(vectorIndex, string(b)) // SINK (LLM08): unsanitized into shared index
	return nil
}
