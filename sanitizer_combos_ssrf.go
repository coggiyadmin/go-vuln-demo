// Combinations #6/#7/#8 — SANITIZER handling × SSRF (CWE-918, Go).
package main

import (
	"net/http"
	"strings"
)

func scEscapeHTML(s string) string { return strings.NewReplacer("<", "_", ">", "_").Replace(s) } // wrong-context
func scFakeSan(s string) string    { return s }                                                  // fake
func scChecked(s string) string { // real wrapper
	if !strings.HasPrefix(s, "https://api.internal/") {
		return "https://api.internal/"
	}
	return s
}

// #6 WRONG-CONTEXT — HTML escape useless for SSRF; should still fire
func scWrongSsrf(w http.ResponseWriter, r *http.Request) { http.Get(scEscapeHTML(r.FormValue("url"))) } // CWE-918

// #7 FAKE sanitizer — no-op; should fire
func scFakeSsrf(w http.ResponseWriter, r *http.Request) { http.Get(scFakeSan(r.FormValue("url"))) } // CWE-918

// #8 CUSTOM-WRAPPER — real allowlist; should NOT fire
func scWrappedSsrf(w http.ResponseWriter, r *http.Request) { http.Get(scChecked(r.FormValue("url"))) } // expect 0
