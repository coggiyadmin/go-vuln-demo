// v09 guard-in-variable
package main
import ("os"; "path/filepath")
var allowed = map[string]bool{"/var/log": true, "/tmp/safe": true}
func read(base, name string) ([]byte, error) {
	if allowed[base] { return os.ReadFile(filepath.Join(base, name)) }
	return nil, nil
}
