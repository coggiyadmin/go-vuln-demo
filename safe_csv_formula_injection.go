// SAFE mirror — csv_formula_injection.go. Prefix with ' so spreadsheet treats cell as text.
package main

import (
	"net/http"
	"os"
)

func safeCsvFormulaExport(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	safe := name
	if safe != "" {
		safe = "'" + safe
	}
	f, _ := os.OpenFile("/var/app/export.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	f.WriteString(safe + ",100\n")
	f.Close()
}
