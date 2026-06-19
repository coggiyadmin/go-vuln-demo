// CWE-331 — Insufficient Entropy (Go). Security token from non-crypto math/rand. NO finding = FN.
package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func ieResetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%06d", rand.Intn(1000000)) // math/rand not a CSPRNG → CWE-331/338
}
