// SAFE mirror — insufficient_entropy.go; uses crypto/rand (CSPRNG). Expect 0 findings.
package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

func safeIeResetToken(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	rand.Read(b) // cryptographically secure
	w.Write([]byte(hex.EncodeToString(b)))
}
