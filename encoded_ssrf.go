// Combination #13 — ENCODED PAYLOAD × SSRF (CWE-918, Go). base64-decode preserves taint.
package main

import (
	"encoding/base64"
	"net/http"
)

func encSsrf(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.FormValue("b64")) // decode
	http.Get(string(raw))                                         // CWE-918
}
