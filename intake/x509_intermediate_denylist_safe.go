// Safe mirror — reject the placeholder untrusted intermediate before Verify.
package intake

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var deniedIntermediates = map[string]struct{}{
	"EXAMPLE-UNTRUSTED-INTERMEDIATE": {},
}

func AcceptChain(pemBytes []byte, roots *x509.CertPool) error {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return errors.New("no certificate")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}
	if _, denied := deniedIntermediates[cert.Subject.CommonName]; denied {
		return errors.New("denylisted intermediate")
	}
	chains, err := cert.Verify(x509.VerifyOptions{Roots: roots})
	if err != nil {
		return err
	}
	for _, chain := range chains {
		for _, c := range chain {
			if _, denied := deniedIntermediates[c.Subject.CommonName]; denied {
				return errors.New("denylisted intermediate")
			}
		}
	}
	return nil
}
