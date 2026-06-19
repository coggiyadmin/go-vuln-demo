// FN probe — CWE-369 Divide By Zero. A user-supplied divisor can be 0.
package main

import (
	"net/http"
	"strconv"
)

func divideAvg(w http.ResponseWriter, r *http.Request) {
	total := 1000
	n, _ := strconv.Atoi(r.URL.Query().Get("n")) // SOURCE — divisor, can be 0
	_ = total / n                                 // divide-by-zero → CWE-369
}
