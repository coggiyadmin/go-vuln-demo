// CWE-807 — Reliance on Untrusted Inputs in a Security Decision (Go). A client header is
// trusted to grant admin access. (Engine gap.) FN probe.
package main

import "net/http"

func usdAdmin(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Is-Admin") == "true" { // trusts attacker-controllable header → CWE-807
		w.Write([]byte("admin panel"))
		return
	}
	http.Error(w, "denied", http.StatusForbidden)
}
