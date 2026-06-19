// FN probe — CWE-1173 Improper Use of Validation Framework. The validator is invoked but its
// result is discarded, so invalid input proceeds. Real vuln; NO finding = FN.
package main

import (
	"net/http"
	"regexp"
)

var improperEmailRe = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

func improperValidate(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	improperEmailRe.MatchString(email) // validation result IGNORED → CWE-1173
	saveImproperSubscriber(email)      // proceeds with unvalidated input
}

func saveImproperSubscriber(_ string) {}
