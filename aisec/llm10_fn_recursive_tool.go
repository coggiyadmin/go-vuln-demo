// Unbounded Consumption FN (OWASP LLM10) — recursive tool re-invokes the model. MISS.
package aisec

import "strings"

func Expand(topic string) string {
	sub := chatSystem("", "Sub-topics of "+topic)
	for _, line := range strings.Split(sub, "\n") {
		if strings.TrimSpace(line) != "" {
			Expand(strings.TrimSpace(line)) // SINK (LLM10 recursive)
		}
	}
	return sub
}
