// SAFE mirror — redos_probe.go; fixed linear-time pattern, user text only.
package main

import (
	"net/http"
	"regexp"
)

var allowedRedos = regexp.MustCompile(`^[a-z0-9_]{1,32}$`)

func safeRedosMatch(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("t")
	allowedRedos.MatchString(text)
}
