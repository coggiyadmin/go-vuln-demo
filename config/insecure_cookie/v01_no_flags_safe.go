package main

import "net/http"

func Set(w http.ResponseWriter, sid string) {
    http.SetCookie(w, &http.Cookie{Name: "SESSIONID", Value: sid, Secure: true, HttpOnly: true, SameSite: http.SameSiteLaxMode})
}
