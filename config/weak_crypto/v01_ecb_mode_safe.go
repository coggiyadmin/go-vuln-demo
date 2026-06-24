package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
)

func EncGCM(key, data []byte) ([]byte, error) {
    b, _ := aes.NewCipher(key)
    gcm, _ := cipher.NewGCM(b)
    nonce := make([]byte, gcm.NonceSize())
    rand.Read(nonce)
    return gcm.Seal(nonce, nonce, data, nil), nil
}
