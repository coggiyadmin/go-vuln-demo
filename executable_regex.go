// FN probe — CWE-624 Executable Regular Expression Error. A user-supplied pattern is compiled
// and drives a replacement, letting the attacker control match logic. NO finding = FN.
package main

import (
	"net/http"
	"regexp"
)

func execRegexRedact(w http.ResponseWriter, r *http.Request) {
	pattern := r.URL.Query().Get("p") // SOURCE — attacker-controlled regex
	text := r.URL.Query().Get("t")
	re, err := regexp.Compile(pattern) // executable user regex → CWE-624 (cf 625)
	if err != nil {
		http.Error(w, "bad pattern", http.StatusBadRequest)
		return
	}
	w.Write([]byte(re.ReplaceAllString(text, "#")))
}
