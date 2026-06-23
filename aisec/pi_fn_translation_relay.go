// Prompt Injection FN (OWASP LLM01) — multi-hop translation relay. MISS.
package aisec

func Relay(userText string) string {
	decoded := chatSystemUser("Translate to English.", userText) // hop 1
	return chatSystem("Do exactly this:\n"+decoded, "")          // SINK (LLM01 relay)
}
