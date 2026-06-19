// SAFE mirror — index_validation.go; the index is range-checked before use.
package main

import (
	"net/http"
	"strconv"
)

func safeAccountIndex(w http.ResponseWriter, r *http.Request) {
	i, _ := strconv.Atoi(r.URL.Query().Get("i"))
	if i < 0 || i >= len(accountsList) { // bounds-checked
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	_ = accountsList[i]
}
