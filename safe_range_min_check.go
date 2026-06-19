// SAFE mirror — range_min_check.go; both lower and upper bounds are checked.
package main

import (
	"net/http"
	"strconv"
)

var safeRangeRecords = []string{"a", "b", "c"}

func safeRangeRecord(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.URL.Query().Get("i"))
	if err == nil && i >= 0 && i < len(safeRangeRecords) { // both bounds checked
		w.Write([]byte(safeRangeRecords[i]))
		return
	}
	http.Error(w, "not found", http.StatusNotFound)
}
