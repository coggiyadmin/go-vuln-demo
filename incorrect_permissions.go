// CWE-276 — Incorrect Default Permissions (Go). A file is created world-writable (0777).
// NO finding = FALSE NEGATIVE. (CWE-732 family.)
package main

import (
	"net/http"
	"os"
)

func ipSave(w http.ResponseWriter, r *http.Request) {
	os.WriteFile("/var/app/data/output.txt", []byte(r.FormValue("data")), 0o777) // world-writable → CWE-276
}
