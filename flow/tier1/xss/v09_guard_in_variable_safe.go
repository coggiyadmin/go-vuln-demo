// v09 guard-in-variable — tag allowlist held in package variable.
package main

var allowedTags = map[string]bool{"p": true, "span": true}

func render(tag, text string) string {
	if allowedTags[tag] {
		return "<" + tag + ">" + text + "</" + tag + ">"
	}
	return "<p>blocked</p>"
}
