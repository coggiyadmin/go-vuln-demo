// SAFE mirror (OWASP LLM07) — no secret in the prompt; key stays server-side.
package aisec

func BuildAgentSystemSafe() string {
	return "You are billing-bot. Use the authorized billing tool for balances."
}
