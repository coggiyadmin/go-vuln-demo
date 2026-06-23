// Unbounded Consumption (OWASP LLM10) — no length/token cap on a user-driven completion. TP.
package aisec

func Summarize(userText string) string { return chatSystem("", userText) } // SINK (LLM10)
