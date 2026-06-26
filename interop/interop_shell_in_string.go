// IL-1 polyglot — Go → shell DSL (CWE-78).
package interop

import (
	"net/http"
	"os/exec"
)

func ShellInString(w http.ResponseWriter, r *http.Request) {
	arg := r.URL.Query().Get("arg") // SOURCE
	// SINK (CWE-78): shell guest-language snippet in a Go string → sh -c.
	cmd := "echo " + arg
	exec.Command("sh", "-c", cmd).Run()
}
