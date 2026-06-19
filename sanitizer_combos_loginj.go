// Combinations #6/#7/#8 — SANITIZER × LOG INJECTION (CWE-117, Go).
package main

import (
	"log"
	"net/http"
	"strings"
)

func scStripHTML(s string) string { return strings.NewReplacer("<", "_", ">", "_").Replace(s) }
func scFakeLog(s string) string   { return s }
func scStripCRLF(s string) string { return strings.Map(func(r rune) bool { return r != '\r' && r != '\n' }, s) }

func scWrongLogInj(w http.ResponseWriter, r *http.Request) {
	log.Println("a " + scStripHTML(r.FormValue("user"))) // CWE-117
}

func scFakeLogInj(w http.ResponseWriter, r *http.Request) {
	log.Println("a " + scFakeLog(r.FormValue("user"))) // CWE-117
}

func scWrappedLogInj(w http.ResponseWriter, r *http.Request) {
	log.Println("a " + scStripCRLF(r.FormValue("user"))) // expect 0
}
