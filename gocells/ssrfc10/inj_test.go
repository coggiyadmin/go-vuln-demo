// Combination #10 — TEST-FILE handling × SSRF (CWE-918, Go). Real sink in a *_test.go
// file. Question: fires by default? does --exclude-tests suppress it?
package ssrfc10

import "net/http"

func handle(r *http.Request) {
	http.Get(r.FormValue("url")) // real CWE-918 in a _test.go file
}
