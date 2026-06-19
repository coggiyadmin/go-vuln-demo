// SAFE mirror — insufficient_logging.go; the privileged action is audit-logged with actor +
// target + outcome.
package main

import (
	"log"
	"net/http"
)

func safeInsufficientDelete(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("userId")
	actor := r.Header.Get("X-Actor")
	doSafeInsufficientDelete(target)
	log.Printf("audit action=delete-user actor=%s target=%s outcome=ok", actor, target)
	w.Write([]byte("deleted"))
}

func doSafeInsufficientDelete(_ string) {}
