// v09 guard-in-variable
package main
import "log"
var allowed = map[string]bool{"INFO": true, "WARN": true}
func logLevel(level, msg string) { if allowed[level] { log.Printf("[%s] %s", level, msg) } }
