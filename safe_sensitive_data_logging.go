// SAFE mirror — sensitive_data_logging.go; only non-sensitive fields are logged.
package main

import (
	"log"
	"net/http"
)

func safeLoginLog(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	_ = r.FormValue("password")            // used for auth, never logged
	log.Printf("login attempt user=%s", user) // no credential in the log
}
