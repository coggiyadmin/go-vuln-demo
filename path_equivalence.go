// CWE-41 — Path Equivalence (Go). An extension check is bypassed by equivalent forms
// (trailing dot/space, './'). NO finding = FALSE NEGATIVE.
package main

import (
	"net/http"
	"os"
	"strings"
)

const peBase = "/srv/app/data/"

func pathEquivGet(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")            // SOURCE
	if strings.HasSuffix(name, ".txt") {   // naive allowlist by extension
		os.ReadFile(peBase + name)         // "secret.txt." / "x/./secret" bypasses → CWE-41
	}
}
