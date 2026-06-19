// CWE-494 — Download of Code Without Integrity Check (Go). A script is fetched and executed
// with no signature/hash verification. (Engine gap — cf #93.) FN probe.
package main

import (
	"io"
	"net/http"
	"os/exec"
)

func dwiSelfUpdate(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get("http://updates.internal/install.sh")
	code, _ := io.ReadAll(resp.Body)
	exec.Command("bash", "-c", string(code)).Run() // executes unverified downloaded code → CWE-494
}
