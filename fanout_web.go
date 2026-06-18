// Combination #11 — FAN-OUT / dedup (Go). One tainted source reaches multiple distinct
// web sinks; expect one finding per sink type.
package main

import (
	"net/http"
	"os"
)

func foWeb(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")                    // single SOURCE
	http.Get(u)                              // sink 1 — ssrf (CWE-918)
	os.ReadFile("/srv/app/data/" + u)        // sink 2 — path_traversal (CWE-22)
	http.Redirect(w, r, u, http.StatusFound) // sink 3 — open_redirect (CWE-601)
}
