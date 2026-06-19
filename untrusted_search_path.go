// CWE-426 — Untrusted Search Path (Go). Bare binary name resolved via $PATH; an
// attacker-controlled PATH entry runs a malicious binary. (Engine gap.) FN probe.
package main

import (
	"net/http"
	"os/exec"
)

func searchPathThumb(w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("src")
	exec.Command("convert", src, "out.png").Run() // 'convert' resolved via $PATH → CWE-426
}
