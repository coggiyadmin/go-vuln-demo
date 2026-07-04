// v09 guard-in-variable
package main
import ("os/exec")
var allowed = map[string]bool{"grep": true, "cat": true}
func run(bin, arg string) error {
	if allowed[bin] { _, err := exec.Command(bin, arg).CombinedOutput(); return err }
	return nil
}
