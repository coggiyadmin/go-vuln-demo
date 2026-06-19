// FN probe — CWE-625 / CWE-1333 ReDoS (user-controlled regex pattern).
// Cross-language mirror of python-vuln-demo/user_regex.py.
package main

import (
	"net/http"
	"regexp"
)

func redosMatch(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("p")
	text := r.URL.Query().Get("t")
	re, _ := regexp.Compile(pattern)
	re.MatchString(text)
}
