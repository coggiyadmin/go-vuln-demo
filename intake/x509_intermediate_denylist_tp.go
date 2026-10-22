// TP — Verify accepts a chain with no intermediate denylist (CWE-295).
// CVE-2026-111943 class: crypto/x509 did not filter the untrusted intermediates.
// Placeholder name only; no certificate material.
package intake

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func AcceptChain(pemBytes []byte, roots *x509.CertPool) error {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return errors.New("no certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	_, err = cert.Verify(x509.VerifyOptions{Roots: roots}) // SINK CWE-295
	return err
}
