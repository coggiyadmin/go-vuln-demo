// FN probe — CWE-749 Exposed Dangerous Method or Function. A privileged maintenance routine is
// exposed on an unauthenticated handler. Real vuln; NO finding = FN.
package main

import (
	"net/http"
	"os/exec"
)

func exposedMaintenance(w http.ResponseWriter, r *http.Request) {
	// privileged maintenance action exposed with no auth gate → CWE-749
	exec.Command("/opt/app/bin/cleanup.sh").Run()
}
