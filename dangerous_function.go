// FN probe — CWE-242 Use of Inherently Dangerous Function. Spawning a shell via `sh -c` is
// dangerous by design (command surface). Real vuln; NO finding = FN.
package main

import (
	"net/http"
	"os/exec"
)

func dangerousCalc(w http.ResponseWriter, r *http.Request) {
	// inherently dangerous: a shell interpreter is invoked → CWE-242
	out, _ := exec.Command("sh", "-c", "expr 1 + 1").Output()
	w.Write(out)
}
