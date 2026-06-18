// Combination #9 — COMMENT / STRING-LITERAL × SSRF (CWE-918, Go). The sink appears only
// in a comment and a string literal — it must NOT fire (0 findings = correct).
package main

import "net/http"

func csSsrf(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	// http.Get(url)   <-- commented sink, must not fire
	example := "http.Get(url)" // string literal, must not fire
	_ = example
	_ = url
}
