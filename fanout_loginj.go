// Combination #11 — FAN-OUT × LOG INJECTION (CWE-117, Go).
package main

import (
	"log"
	"net/http"
)

func fanoutLogInjHandler(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	log.Println("a " + u)
	log.Println("b " + u)
	log.Println("c " + u)
}
