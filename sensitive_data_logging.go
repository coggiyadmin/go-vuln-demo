// FN probe — CWE-779 Logging of Sensitive Data. The user's password is written to the log.
// (CWE-532 family.)
package main

import (
	"log"
	"net/http"
)

func loginLog(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	pw := r.FormValue("password") // SOURCE — credential
	log.Printf("login attempt user=%s password=%s", user, pw) // logs the password → CWE-779
}
