// SAFE mirror — improper_validation.go; the validator result is enforced before use.
package main

import (
	"net/http"
	"regexp"
)

var safeImproperEmailRe = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

func safeImproperValidate(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if !safeImproperEmailRe.MatchString(email) { // result enforced
		http.Error(w, "invalid email", http.StatusBadRequest)
		return
	}
	saveSafeImproperSubscriber(email)
}

func saveSafeImproperSubscriber(_ string) {}
