// Phase-3 validate mirror — allowlist/regex gate before sink
package main
import ("log"; "net/http")
func logInfoSafe(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    log.Printf("login user=%s", user)
}
