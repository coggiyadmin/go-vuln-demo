// SAFE mirror — executable_regex.go; a fixed server-side pattern redacts user text only.
package main

import (
	"net/http"
	"regexp"
)

var execDigitsRe = regexp.MustCompile(`\d`)

func safeExecRegexRedact(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("t")
	w.Write([]byte(execDigitsRe.ReplaceAllString(text, "#"))) // static pattern
}
