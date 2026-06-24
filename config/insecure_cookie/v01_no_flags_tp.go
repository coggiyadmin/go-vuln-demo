package main

import "net/http"

func Set(w http.ResponseWriter, sid string) {
    http.SetCookie(w, &http.Cookie{Name: "SESSIONID", Value: sid, Secure: false, HttpOnly: false}) // SINK CWE-614
}
