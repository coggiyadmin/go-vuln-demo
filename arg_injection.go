// CWE-88 — Argument Injection (Go). User input as a CLI arg without `--`. NO finding = FN.
package main

import (
	"net/http"
	"os/exec"
)

func argInjLog(w http.ResponseWriter, r *http.Request) {
	branch := r.FormValue("branch")          // SOURCE
	exec.Command("git", "log", branch).Run() // branch like "--output=..." → CWE-88
}
