// FN probe — CWE-367 Time-of-check Time-of-use (TOCTOU) race condition.
// File is checked then opened separately; attacker can swap it in the window.
package main

import (
	"net/http"
	"os"
)

const toctouBase = "/var/app/data/"

func toctouWrite(w http.ResponseWriter, r *http.Request) {
	path := toctouBase + r.URL.Query().Get("f")
	if _, err := os.Stat(path); err == nil { // CHECK
		f, _ := os.OpenFile(path, os.O_WRONLY, 0o600) // USE — race window → CWE-367
		defer f.Close()
		f.WriteString(r.URL.Query().Get("d"))
	}
}
