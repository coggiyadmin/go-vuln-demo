// SAFE mirror — sensitive_info_exposure.go; secret written to a private, non-served path.
package main

import (
	"net/http"
	"os"
)

func safeExportConfig(w http.ResponseWriter, r *http.Request) {
	// not web-accessible, owner-only permissions
	os.WriteFile("/var/app/private/config.txt", []byte("API_KEY="+os.Getenv("API_KEY")), 0o600)
}
