// CWE-250 — Execution with Unnecessary Privileges (Go). The process escalates to root before
// an operation that does not require it. (Engine gap.) FN probe.
package main

import (
	"net/http"
	"os/exec"
	"syscall"
)

func upClearCache(w http.ResponseWriter, r *http.Request) {
	syscall.Setuid(0)                                  // escalate to root unnecessarily → CWE-250
	exec.Command("rm", "-rf", "/var/app/cache").Run()
}
