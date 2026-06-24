// FP-target — constant log line must not be flagged log injection.
package fpcorpus
import "log"
func Start() { log.Print("service started") } // constant — NOT a sink
