// CWE-470 — Unsafe Reflection (Go). A plugin (.so) is loaded from a user-controlled path,
// executing arbitrary code. NO finding = FALSE NEGATIVE.
package main

import (
	"net/http"
	"plugin"
)

func reflectLoad(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("mod") // SOURCE
	plugin.Open(path)          // loads arbitrary plugin code → CWE-470
}
