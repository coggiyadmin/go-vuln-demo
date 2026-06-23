// System-Prompt Leakage FN (OWASP LLM07) — secret leaks inside an error string. MISS.
package aisec

import (
	"fmt"
	"os"
)

func Run() error {
	return fmt.Errorf("config dump: system=triage key=%s", os.Getenv("BILLING_API_KEY")) // SINK (indirect)
}
