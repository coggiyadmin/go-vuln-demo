// Prompt Injection DIRECT (OWASP LLM01).
package aisec

import "fmt"

func BuildPrompt(user string) string {
	return "You are a support bot. Follow policy.\n" + user
}

func Answer(user string) string {
	prompt := BuildPrompt(user)
	return fmt.Sprintf("llm:%s", prompt)
}
