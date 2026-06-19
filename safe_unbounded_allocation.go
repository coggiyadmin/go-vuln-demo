// SAFE mirror — unbounded_allocation.go; requested size clamped to a hard limit.
package main

import (
	"net/http"
	"strconv"
)

const maxAllocBytes = 1 << 20 // 1 MiB cap

func safeAllocBuf(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))
	if n < 0 || n > maxAllocBytes { // bounded
		http.Error(w, "too large", http.StatusBadRequest)
		return
	}
	buf := make([]byte, n)
	_ = buf
}
