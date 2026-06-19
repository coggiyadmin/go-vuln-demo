// SAFE mirror — static_iv.go; a fresh random IV is generated per encryption. Expect 0.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"net/http"
)

var safeSivKey = []byte("0123456789abcdef0123456789abcdef")

func safeSivEnc(w http.ResponseWriter, r *http.Request) {
	block, _ := aes.NewCipher(safeSivKey)
	iv := make([]byte, 16)
	rand.Read(iv) // unique IV per encryption
	data := []byte(r.FormValue("d"))
	enc := make([]byte, len(data))
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(enc, data)
	w.Write(append(iv, enc...))
}
