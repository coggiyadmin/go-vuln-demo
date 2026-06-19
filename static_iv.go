// CWE-1204 — Weak/Static Initialization Vector (Go). A static IV is used for AES-CFB. NO finding = FN.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"net/http"
)

var sivKey = []byte("0123456789abcdef0123456789abcdef")

func sivEnc(w http.ResponseWriter, r *http.Request) {
	block, _ := aes.NewCipher(sivKey)
	iv := make([]byte, 16) // static all-zero IV → CWE-1204
	data := []byte(r.FormValue("d"))
	enc := make([]byte, len(data))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(enc, data)
	w.Write(enc)
}
