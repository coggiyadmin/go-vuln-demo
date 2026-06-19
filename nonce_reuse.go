// CWE-323 — Reusing a Nonce in Encryption (Go). A fixed AES-GCM nonce is reused across
// encryptions. Real vuln; NO finding = FALSE NEGATIVE.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"net/http"
)

var nrKey = []byte("0123456789abcdef0123456789abcdef")
var nrNonce = make([]byte, 12) // static nonce, reused → CWE-323

func nrEnc(w http.ResponseWriter, r *http.Request) {
	block, _ := aes.NewCipher(nrKey)
	gcm, _ := cipher.NewGCM(block)
	enc := gcm.Seal(nil, nrNonce, []byte(r.FormValue("d")), nil) // same nonce every call → CWE-323
	w.Write(enc)
}
