// TP anchor (CWE-78) for #170 — a REAL OS command exec of untrusted HTTP input. Proves the engine
// fires command_injection on a genuine OS sink, so safe_redis_protocol.go (a RESP protocol verb)
// staying clean is a credited distinction, not an engine blind spot.
package libapi

import (
	"net/http"
	"os/exec"
)

// RunUnsafe execs an attacker-controlled query parameter.
func RunUnsafe(w http.ResponseWriter, r *http.Request) {
	exec.Command("sh", "-c", r.URL.Query().Get("cmd")).Run() // SINK — real OS command execution
}
