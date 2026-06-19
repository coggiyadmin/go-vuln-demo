// FN probe — CWE-1284 Improper Validation of Specified Quantity in Input. A user-supplied
// order quantity is used unchecked (negative / huge) in allocation and pricing. NO finding = FN.
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

const unitPrice1284 = 1000

func quantityOrder(w http.ResponseWriter, r *http.Request) {
	qty, _ := strconv.Atoi(r.URL.Query().Get("qty")) // SOURCE — unvalidated quantity
	slots := make([]int, qty)                         // negative/huge used unchecked → CWE-1284
	fmt.Fprintf(w, "%d", unitPrice1284*qty+len(slots))
}
