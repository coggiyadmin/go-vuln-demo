package main

import "crypto/sha256"

func HashSHA256(data []byte) [32]byte { return sha256.Sum256(data) }
