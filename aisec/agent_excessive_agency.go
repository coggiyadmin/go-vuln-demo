// Excessive Agency (OWASP LLM06).
package aisec

import "os/exec"

func ShellTool(command string) ([]byte, error) {
	return exec.Command("sh", "-c", command).CombinedOutput()
}
