// v10 allowlist-set — html.EscapeString after map membership gate.
package main

import ("html"; "strings")

var safeNames = map[string]bool{"alice": true, "bob": true}

func greet(name string) string {
	if safeNames[name] {
		return "<p>Hello " + html.EscapeString(name) + "</p>"
	}
	return "<p>unknown</p>"
}
