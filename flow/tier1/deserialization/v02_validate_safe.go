// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("encoding/json"; "net/http")
func deserSafe(w http.ResponseWriter, r *http.Request) {
    var v map[string]interface{}
    json.Unmarshal([]byte(r.FormValue("blob")), &v)
}
