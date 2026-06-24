// TP (CWE-614/1004, go) — cookie set without Secure/HttpOnly flags.
package fpcorpus

import "net/http"

// SetSession writes a session cookie missing the Secure and HttpOnly attributes.
func SetSession(w http.ResponseWriter, sid string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "SESSIONID",
		Value:    sid,
		Secure:   false, // SINK (CWE-614): transmitted over plaintext
		HttpOnly: false, // SINK (CWE-1004): readable from JS
	})
}
