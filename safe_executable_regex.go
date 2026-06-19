// SAFE mirror — executable_regex.go; a fixed server-side pattern redacts user text only.
package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var execDigitsRe = regexp.MustCompile(`\d`)

func safeExecRegexRedact(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("t")
	n := len(execDigitsRe.FindAllString(text, -1)) // static pattern; no user text reflected
	fmt.Fprintf(w, "%d", n)
}
