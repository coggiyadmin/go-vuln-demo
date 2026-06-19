// CWE-472 — External Control of Assumed-Immutable Web Parameter (Go). Server trusts a
// client-supplied 'price' for a financial decision. (Engine gap.) FN probe.
package main

import (
	"net/http"
	"strconv"
)

func paramTamperCheckout(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64) // SOURCE — client-controlled
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	_ = price * float64(qty) // trusts client-set price → CWE-472
}
