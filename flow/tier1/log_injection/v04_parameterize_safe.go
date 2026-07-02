// Phase-3 parameterize mirror — bound params / prepared API
package main
import ("log"; "net/http")
func logInfoSafe(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    log.Printf("login user=%s", user)
}
