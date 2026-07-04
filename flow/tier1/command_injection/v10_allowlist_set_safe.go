// v10 allowlist-set
package main
import "os/exec"
var safePairs = map[string]bool{"grep:/var/log/app.log": true}
func run(cmd, path string) error {
	if safePairs[cmd+":"+path] { _, err := exec.Command(cmd, path).CombinedOutput(); return err }
	return nil
}
