// v09 guard-in-variable
package main
var allowed = map[string]bool{"name": true, "status": true}
func find(col, val string) map[string]string {
	if allowed[col] { return map[string]string{col: val} }
	return nil
}
