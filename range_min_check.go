// FN probe — CWE-839 Numeric Range Comparison Without Minimum Check. The index is checked
// against the upper bound only, so a negative offset is accepted. NO finding = FN.
package main

import (
	"net/http"
	"strconv"
)

var rangeRecords = []string{"a", "b", "c"}

func rangeRecord(w http.ResponseWriter, r *http.Request) {
	i, _ := strconv.Atoi(r.URL.Query().Get("i")) // SOURCE
	if i < len(rangeRecords) {                    // upper bound only — no `i >= 0` → CWE-839
		w.Write([]byte(rangeRecords[i]))
		return
	}
	http.Error(w, "not found", http.StatusNotFound)
}
