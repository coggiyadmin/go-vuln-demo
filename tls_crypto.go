// DEMO FILE — patterns from famous SAST benchmarks (gosec) not yet exercised.
//
// Aligns with gosec: G402 (TLS InsecureSkipVerify), G404 (weak random),
// G401/G501 (MD5), G505 (SHA1). CWE-295, CWE-327, CWE-328, CWE-330.
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
)

// CWE-295 / gosec G402 — TLS certificate verification disabled
func insecureClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // skips cert validation → MITM
	}
	return &http.Client{Transport: tr}
}

// CWE-328 / gosec G401 — MD5 for security
func hashPassword(pw string) string {
	sum := md5.Sum([]byte(pw)) // weak hash
	return fmt.Sprintf("%x", sum)
}

// CWE-327 / gosec G505 — SHA1
func fingerprint(data []byte) string {
	sum := sha1.Sum(data) // weak hash
	return fmt.Sprintf("%x", sum)
}

// CWE-330 / gosec G404 — non-crypto RNG for a security token
func resetToken() int {
	return rand.Intn(1000000) // predictable
}
