// FN probe — CWE-778 Insufficient Logging. A privileged user deletion is performed with no
// audit log, so attacks are untraceable. NO finding = FN.
package main

import "net/http"

func insufficientDelete(w http.ResponseWriter, r *http.Request) {
	target := r.FormValue("userId")
	// SECURITY EVENT (privileged delete) performed with NO audit log → CWE-778
	doInsufficientDelete(target)
	w.Write([]byte("deleted"))
}

func doInsufficientDelete(_ string) {}
