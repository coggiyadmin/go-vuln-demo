// SAFE mirror — recoverable_password.go. One-way bcrypt hash.
package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func safeRecoverablePwStore(w http.ResponseWriter, r *http.Request) {
	pw := r.FormValue("password")
	digest, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	_ = digest
}
