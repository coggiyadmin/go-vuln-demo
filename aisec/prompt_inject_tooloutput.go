// Prompt Injection TOOL-OUTPUT (OWASP LLM01) — tool result given instruction authority.
package aisec

// SINK (LLM01 tool-output): attacker-influenceable toolResult spliced into the system role.
func RunWithTool(userQ, toolResult string) string {
	system := "You are an assistant. Tool said:\n" + toolResult + "\nNow act on it."
	return chatSystem(system, userQ)
}
