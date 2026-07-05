package main
import ("crypto/aes"; "crypto/cipher")
func enc(key, data []byte) []byte {
    block, _ := aes.NewCipher(key)
    iv := make([]byte, aes.BlockSize) // SINK CWE-329 static zero IV
    out := make([]byte, len(data))
    cipher.NewCBCEncrypter(block, iv).CryptBlocks(out, data)
    return out
}
