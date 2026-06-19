// CWE-261 — Weak Encoding for Password (Go). base64 used as if it protected the password.
// (Engine gap.) FN probe.
package main

import (
	"encoding/base64"
	"net/http"
	"os"
)

func weakPwEncode(w http.ResponseWriter, r *http.Request) {
	enc := base64.StdEncoding.EncodeToString([]byte(r.FormValue("password"))) // encoding ≠ protection → CWE-261
	os.WriteFile("/var/app/pw.txt", []byte(enc), 0o644)
}
