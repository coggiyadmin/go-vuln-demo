// v09 guard-in-variable
package main
import "encoding/json"
var allowed = map[string]bool{"{}": true}
func loads(raw string) (map[string]interface{}, error) {
	if allowed[raw] { var m map[string]interface{}; err := json.Unmarshal([]byte(raw), &m); return m, err }
	return nil, nil
}
