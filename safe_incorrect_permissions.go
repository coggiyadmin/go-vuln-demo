// SAFE mirror — incorrect_permissions.go; owner-only permissions (0600). Expect 0 findings.
package main

import (
	"net/http"
	"os"
)

func safeIpSave(w http.ResponseWriter, r *http.Request) {
	os.WriteFile("/var/app/data/output.txt", []byte(r.FormValue("data")), 0o600) // owner read/write only
}
