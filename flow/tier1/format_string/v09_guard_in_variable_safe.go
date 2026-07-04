// v09 guard-in-variable
package main
import "fmt"
var allowed = map[string]bool{"user=%s": true}
func format(t, v string) string { if allowed[t] { return fmt.Sprintf(t, v) }; return "" }
