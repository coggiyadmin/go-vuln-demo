// FP-target (#128/#140) — LIBRARY profile. Public utility; cmd is caller-supplied.
package profile

import "os/exec"

func Run(cmd string) ([]byte, error) {
	return exec.Command("sh", "-c", cmd).Output() // caller-supplied, not attacker source
}
