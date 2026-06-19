// FN probe — CWE-497 Exposure of Sensitive System Information.
// A debug endpoint returns the full environment and platform details to the client.
package main

import (
	"net/http"
	"os"
	"runtime"
	"strings"
)

func debugInfo(w http.ResponseWriter, r *http.Request) {
	// leaks env vars (secrets), versions, paths to an unauthorized actor → CWE-497
	w.Write([]byte(strings.Join(os.Environ(), "\n") + runtime.GOOS + runtime.Version()))
}
