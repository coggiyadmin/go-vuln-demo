// v09 guard-in-variable
package main
var allowed = map[string]bool{"text/plain": true}
func setCt(v string) string { if allowed[v] { return "Content-Type: " + v }; return "" }
