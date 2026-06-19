// CWE-73 — External Control of File Name or Path (Go). NO finding = FN.
package main

import (
	"net/http"
	"os"
)

func extFileRead(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path") // SOURCE — full path
	os.ReadFile(path)           // arbitrary file read → CWE-73
}
