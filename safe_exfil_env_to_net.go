// SAFE mirror — the env secret is used purely LOCALLY (to sign a record); nothing
// sensitive crosses a network boundary, and there is no dynamic header. ZERO exfil.
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func signRecord(record string) string {
	key := os.Getenv("INTERNAL_API_KEY") // local use only — never sent anywhere
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(record))
	return hex.EncodeToString(m.Sum(nil))
}
