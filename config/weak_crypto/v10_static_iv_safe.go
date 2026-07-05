package main
import ("crypto/aes"; "crypto/cipher"; "crypto/rand"; "io")
func enc(key, data []byte) []byte {
    block, _ := aes.NewCipher(key)
    iv := make([]byte, aes.BlockSize)
    io.ReadFull(rand.Reader, iv)
    out := make([]byte, len(data))
    cipher.NewCBCEncrypter(block, iv).CryptBlocks(out, data)
    return out
}
