// Combination #5 — OOP OBJECT FLOW × LOG INJECTION (CWE-117, Go). Taint stored in a
// struct field, written to a log sink. NO finding = FALSE NEGATIVE.
package main

import (
	"log"
	"net/http"
)

type auditRec struct{ actor string }

func (a auditRec) recordDirect() {
	log.Println("login by " + a.actor) // log sink (CWE-117)
}

func logInjOOPHandler(w http.ResponseWriter, r *http.Request) {
	a := auditRec{actor: r.FormValue("user")} // SOURCE -> struct field
	a.recordDirect()
}
