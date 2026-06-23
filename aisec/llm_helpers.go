// Test helpers — stand-in LLM client so aisec probes compile without a live SDK.
package aisec

func chatSystem(system, user string) string     { _ = system; _ = user; return "" }
func chatSystemUser(system, user string) string { _ = system; _ = user; return "" }
