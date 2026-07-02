// SAFE — fixed argv only.
package interop

import "os/exec"

func RunSafe() error {
    return exec.Command("echo", "ok").Run()
}
