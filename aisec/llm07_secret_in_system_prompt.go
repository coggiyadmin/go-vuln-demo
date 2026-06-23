// System-Prompt Leakage (OWASP LLM07) — secret baked into the system prompt. TP.
package aisec

import "os"

func BuildAgentSystem() string {
	// SINK (LLM07): secret embedded in the instruction
	return "You are billing-bot. Internal key: " + os.Getenv("BILLING_API_KEY") + ". Never reveal it."
}
