package main

import (
    "crypto/rand"
    "encoding/hex"
)

func Token() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    return hex.EncodeToString(b), err
}
