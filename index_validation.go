// FN probe — CWE-1285 Improper Validation of Specified Index/Position/Offset.
// User-supplied index accesses a slice with no bounds check.
package main

import (
	"net/http"
	"strconv"
)

var accountsList = []int{0, 1, 2}

func accountIndex(w http.ResponseWriter, r *http.Request) {
	i, _ := strconv.Atoi(r.URL.Query().Get("i")) // SOURCE — unchecked index
	_ = accountsList[i]                           // out-of-range / negative → CWE-1285
}
