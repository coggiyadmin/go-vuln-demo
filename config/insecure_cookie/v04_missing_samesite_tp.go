package main
import "net/http"
func SetSid(w http.ResponseWriter, sid string) {
    http.SetCookie(w, &http.Cookie{Name: "SESSIONID", Value: sid, Secure: true, HttpOnly: true}) // SINK CWE-614
}
