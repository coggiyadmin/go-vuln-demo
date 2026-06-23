// SAFE mirror (OWASP LLM10) — input length capped before the call.
package aisec

import "errors"

const maxInput = 8000

func SummarizeSafe(userText string) (string, error) {
	if len(userText) > maxInput {
		return "", errors.New("input too large")
	}
	return chatSystem("", userText), nil
}
