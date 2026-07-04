// v09 guard-in-variable
package main
var allowed = map[string]bool{"1+1": true}
func evalExpr(expr string) string { if allowed[expr] { return expr }; return "" }
