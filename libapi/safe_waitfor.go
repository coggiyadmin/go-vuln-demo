// TN — fixed argv exec (#124).
package libapi

import "os/exec"

func WaitFixed() error {
	return exec.Command("echo", "ok").Run()
}
