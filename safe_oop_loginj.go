// SAFE mirror — oop_loginj.go; CR/LF stripped before logging. Expect 0 findings.
package main

import (
	"log"
	"net/http"
	"strings"
)

type safeAuditRec struct{ actor string }

func (a safeAuditRec) recordDirect() {
	clean := strings.NewReplacer("\r", "", "\n", "").Replace(a.actor) // strip CRLF
	log.Println("login by " + clean)
}

func safeLogInjOOPHandler(w http.ResponseWriter, r *http.Request) {
	a := safeAuditRec{actor: r.FormValue("user")}
	a.recordDirect()
}
