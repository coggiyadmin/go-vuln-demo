// CWE-335 — Incorrect Usage of Seeds in PRNG (Go). The PRNG is seeded with the current time.
// Real vuln; NO finding = FALSE NEGATIVE.
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func psSessionID(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix()) // predictable seed → CWE-335
	fmt.Fprintf(w, "%d", rand.Int63())
}
