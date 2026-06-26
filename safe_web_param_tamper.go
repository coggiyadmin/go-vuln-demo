// SAFE mirror — web_param_tamper.go. Price from server catalog, not client field.
package main

import "net/http"

var catalog = map[string]float64{"sku-1": 9.99}

func safeWebParamCheckout(w http.ResponseWriter, r *http.Request) {
	sku := r.FormValue("sku")
	qty := 1
	_ = qty
	price := catalog[sku]
	_ = price
}
