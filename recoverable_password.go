// CWE-257 — Storing Passwords in a Recoverable Format (Go). AES-encrypted (reversible) with
// a key on hand. (Engine gap.) FN probe.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"net/http"
	"os"
)

var recovKey = make([]byte, 32)

func recoverablePwStore(w http.ResponseWriter, r *http.Request) {
	block, _ := aes.NewCipher(recovKey)
	pw := []byte(r.FormValue("password")) // SOURCE
	enc := make([]byte, len(pw))
	iv := make([]byte, 16)
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(enc, pw) // reversible → CWE-257
	os.WriteFile("/var/app/pw.bin", enc, 0o644)
}
