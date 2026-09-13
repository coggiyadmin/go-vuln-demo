// SAFE mirror — sensitive_data_logging.go; only non-sensitive fields are logged, CR/LF
// stripped so the entry cannot be forged (no CWE-117 either).
package main

import (
	"log"
	"net/http"
	"strings"
)

func safeLoginLog(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	_ = r.FormValue("password")                                  // used for auth, never logged
	safeUser := strings.NewReplacer("\r", "", "\n", "").Replace(user)
	log.Printf("login attempt user=%s", safeUser) // no credential in the log
}
