// FN probe — CWE-770 Allocation of Resources Without Limits or Throttling.
// Buffer allocated from a user-supplied size with no upper bound → memory-exhaustion DoS.
package main

import (
	"net/http"
	"strconv"
)

func allocBuf(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n")) // SOURCE — attacker-controlled size
	buf := make([]byte, n)                        // unbounded allocation → CWE-770
	_ = buf
}
