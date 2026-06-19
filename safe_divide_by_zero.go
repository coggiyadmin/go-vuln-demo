// SAFE mirror — divide_by_zero.go; the divisor is checked for zero.
package main

import (
	"net/http"
	"strconv"
)

func safeDivideAvg(w http.ResponseWriter, r *http.Request) {
	total := 1000
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	if n == 0 { // guarded
		http.Error(w, "n must be non-zero", http.StatusBadRequest)
		return
	}
	_ = total / n
}
