// SAFE mirror — quantity_validation.go; the quantity is validated against a sane range.
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const (
	safeUnitPrice1284 = 1000
	maxQty1284        = 100
)

func safeQuantityOrder(w http.ResponseWriter, r *http.Request) {
	qty, err := strconv.Atoi(r.URL.Query().Get("qty"))
	if err != nil || qty < 1 || qty > maxQty1284 { // validated range
		http.Error(w, "invalid quantity", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%d", safeUnitPrice1284*qty)
}
