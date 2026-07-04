// v09 guard-in-variable
package main
var allowed = map[string]bool{"//users/user[@id='1']": true}
func query(expr string) string { if allowed[expr] { return expr }; return "" }
