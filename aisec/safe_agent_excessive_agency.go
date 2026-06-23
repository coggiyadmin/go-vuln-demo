// SAFE mirror — fixed argv, no shell.
package aisec

import "os/exec"

var allow = map[string]bool{"pwd": true, "date": true}

func ShellTool(command string) ([]byte, error) {
	if !allow[command] {
		return nil, exec.ErrNotFound
	}
	return exec.Command(command).CombinedOutput()
}
