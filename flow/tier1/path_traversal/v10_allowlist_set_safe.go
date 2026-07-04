// v10 allowlist-set
package main
import "os"
var safeFiles = map[string]bool{"/var/log/app.log": true}
func read(target string) ([]byte, error) {
	if safeFiles[target] { return os.ReadFile(target) }
	return nil, nil
}
