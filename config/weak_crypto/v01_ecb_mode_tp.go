// weak_crypto TP — AES ECB (CWE-327).
package main

import "crypto/aes"

func EncECB(key, data []byte) []byte {
    c, _ := aes.NewCipher(key)
    out := make([]byte, len(data))
    c.Encrypt(out, data) // ECB — no IV → CWE-327
    return out
}
