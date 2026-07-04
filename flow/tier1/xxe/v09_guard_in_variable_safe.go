// v09 guard-in-variable
package main
var allowed = map[string]bool{"<root/>": true}
func parse(raw string) string { if allowed[raw] { return raw }; return "" }
