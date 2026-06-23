// SAFE mirror (OWASP LLM10) — bounded by a hard max-steps budget.
package aisec

const maxSteps = 8

func AgentSafe(goal string) []string {
	history := []string{goal}
	for i := 0; i < maxSteps; i++ {
		msg := chatSystem("", joinHistorySafe(history))
		history = append(history, msg)
		if containsSafe(msg, "DONE") {
			break
		}
	}
	return history
}

func joinHistorySafe(h []string) string { return "" }
func containsSafe(s, sub string) bool   { return false }
