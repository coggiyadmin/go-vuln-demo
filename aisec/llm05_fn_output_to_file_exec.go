// Improper Output Handling FN (OWASP LLM05) — deferred exec; output written then run later. MISS.
package aisec

import (
	"os"
	"os/exec"
)

const script = "/var/app/plugins/generated.sh"

func Stage(code string) error { return os.WriteFile(script, []byte(code), 0o755) } // SOURCE
func Activate() ([]byte, error) { return exec.Command("sh", script).Output() }     // SINK (deferred)
