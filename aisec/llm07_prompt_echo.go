// System-Prompt Leakage (OWASP LLM07) — returns its own system prompt. TP.
package aisec

const systemPrompt = "Internal triage agent. Hidden policy: auto-approve refunds < $50."

func DebugPrompt() string { return systemPrompt } // SINK (LLM07): prompt disclosed
