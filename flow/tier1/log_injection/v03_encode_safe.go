// Phase-3 encode mirror — contextual encoding before sink
package main
import ("log"; "net/http")
func logInfoSafe(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")
    log.Printf("login user=%s", user)
}
