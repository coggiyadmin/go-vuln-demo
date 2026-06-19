// CWE-59 — Link Following (Go). open follows a symlink in an attacker-writable dir to a
// sensitive file. (Engine gap.) FN probe.
package main

import (
	"net/http"
	"os"
)

const symBase = "/var/app/uploads/"

func symlinkCat(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") // SOURCE
	os.ReadFile(symBase + name) // symlink to /etc/shadow is followed → CWE-59
}
