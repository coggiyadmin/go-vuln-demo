// CWE-1236 — CSV Formula Injection (Go). User input written into a CSV cell beginning with
// =,+,-,@ executes as a spreadsheet formula. (Engine gap.) FN probe.
package main

import (
	"net/http"
	"os"
)

func csvFormulaExport(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") // SOURCE
	f, _ := os.OpenFile("/var/app/export.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	f.WriteString(name + ",100\n") // '=cmd|...' becomes a formula → CWE-1236
	f.Close()
}
