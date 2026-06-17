// NEGATIVE TEST FILE — safe mirror of supplychain.go DNS C2 pattern.
//
// The scanner MUST produce ZERO security findings here. Any finding is a
// FALSE POSITIVE.
package main

func init() {
	// SAFE init — no DNS polling, no remote command execution.
}
