// TP (CWE-117 / cognium-dev#107) — a tainted HTTP query value logged via log.Printf. The Go
// log_injection sink was added in circle-ir 3.73.0; with an entry-point-reachable source the sink
// must fire. (The prior bare-param form was not entry-point-reachable, so it under-tested the fix.)
package fpcorpus

import (
	"log"
	"net/http"
)

// OnLogin logs an attacker-controlled request parameter.
func OnLogin(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user") // attacker-controlled HTTP source
	log.Printf("login user=%s", user) // SINK (CWE-117 log injection)
}
