// Unbounded Consumption (OWASP LLM10) — agent loop with no iteration cap. TP.
package aisec

func Agent(goal string) []string {
	history := []string{goal}
	for { // SINK (LLM10): no max-steps guard
		msg := chatSystem("", joinHistory(history))
		history = append(history, msg)
		if contains(msg, "DONE") {
			return history
		}
	}
}

func joinHistory(h []string) string { return "" }
func contains(s, sub string) bool   { return false }
