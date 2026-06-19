// FN probe — CWE-183 Permissive List of Allowed Inputs. The host allow-check uses substring
// containment, so `evil-trusted.com.attacker.net` passes. Real vuln; NO finding = FN.
package main

import (
	"net/http"
	"strings"
)

func permissiveFetch(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")     // SOURCE
	if strings.Contains(target, "trusted.com") { // overly-permissive substring → CWE-183
		http.Redirect(w, r, target, http.StatusFound)
	}
}
