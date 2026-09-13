// SAFE mirror — insufficient_logging.go; the privileged action is audit-logged with actor +
// target + outcome, CR/LF stripped so the log line cannot be forged (no CWE-117 either).
package main

import (
	"log"
	"net/http"
	"strings"
)

var logLineSanitizer = strings.NewReplacer("\r", "", "\n", "")

func safeInsufficientDelete(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("userId")
	actor := r.Header.Get("X-Actor")
	safeActor := logLineSanitizer.Replace(actor)
	safeTarget := logLineSanitizer.Replace(target)
	doSafeInsufficientDelete(target)
	log.Printf("audit action=delete-user actor=%s target=%s outcome=ok", safeActor, safeTarget)
	w.Write([]byte("deleted"))
}

func doSafeInsufficientDelete(_ string) {}
