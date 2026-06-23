// Improper Output Handling (OWASP LLM05) — model output run as a shell command. TP.
package aisec

import "os/exec"

func RunGenerated(modelOutput string) ([]byte, error) {
	// SINK (LLM05): model output executed by the shell
	return exec.Command("sh", "-c", modelOutput).Output()
}
