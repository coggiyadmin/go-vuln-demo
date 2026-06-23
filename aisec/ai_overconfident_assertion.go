// Provenance/Misinfo (OWASP LLM09) — fabricated API usage (hallucinated method). TP.
package aisec

import "crypto/sha256"

func SecureToken(seed string) []byte {
	h := sha256.New()
	h.Write([]byte(seed))
	// SINK (LLM09): sha256.Hash has no SecureSum — invented API presented as real
	return h.SecureSum(nil)
}
