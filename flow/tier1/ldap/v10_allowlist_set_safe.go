// v09 guard-in-variable
package main
var allowed = map[string]bool{"dc=example,dc=com": true}
func search(base, filt string) string { if allowed[base] { return base + filt }; return "" }
