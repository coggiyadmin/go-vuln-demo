// SAFE mirror — arg_injection.go; `--` terminates option parsing. Expect 0 findings.
package main

import (
	"net/http"
	"os/exec"
)

func safeArgInjLog(w http.ResponseWriter, r *http.Request) {
	branch := r.FormValue("branch")
	exec.Command("git", "log", "--", branch).Run() // branch can't be a flag
}
