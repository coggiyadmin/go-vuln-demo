// Combination #13 — ENCODED × LOG INJECTION (CWE-117, Go).
package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
)

func encodedLogInjB64Handler(w http.ResponseWriter, r *http.Request) {
	u, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	log.Println("a " + string(u))
}

func encodedLogInjURLHandler(w http.ResponseWriter, r *http.Request) {
	u, _ := url.QueryUnescape(r.FormValue("d"))
	log.Println("a " + u)
}
