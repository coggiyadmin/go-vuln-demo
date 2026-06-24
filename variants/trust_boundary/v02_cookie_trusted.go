// trust_boundary variant: trusted cookie from form input.
package main
import ("net/http")
func trustCookie(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{Name: "is_admin", Value: r.FormValue("admin")}) // SINK CWE-501
}
