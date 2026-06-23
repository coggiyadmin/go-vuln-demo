// SAFE mirror (OWASP LLM05) — model output parsed as an integer, never executed.
package aisec

import "strconv"

func RunGenerated(modelOutput string) (int, error) {
	return strconv.Atoi(modelOutput)
}
