// CWE-916 — Weak Password Hash (Go). Fast, unsalted SHA-256 for passwords. NO finding = FN.
package main

import (
	"crypto/sha256"
	"net/http"
)

func weakPwHash(w http.ResponseWriter, r *http.Request) {
	pw := r.FormValue("password")  // SOURCE
	_ = sha256.Sum256([]byte(pw))  // fast unsalted hash → CWE-916
}
