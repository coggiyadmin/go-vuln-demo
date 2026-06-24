// TP (CWE-117 / cognium-dev#107) — log.Printf with a tainted arg; the Go log_injection sink.
package fpcorpus
import "log"
func OnLogin(user string) { log.Printf("login user=%s", user) } // SINK (engine #107 FN target)
