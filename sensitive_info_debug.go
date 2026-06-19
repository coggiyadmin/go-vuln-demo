// FN probe — CWE-215 Insertion of Sensitive Information Into Debugging Code. A debug branch
// returns the error detail and DB credentials to the client. NO finding = FN.
package main

import (
	"fmt"
	"net/http"
	"os"
)

var debug215 = os.Getenv("DEBUG") == "1"

func sensitiveDebug(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("boom")
	if debug215 {
		// leaks error detail + secrets through a debug path → CWE-215
		fmt.Fprintf(w, "%v\nDB_PASSWORD=%s", err, os.Getenv("DB_PASSWORD"))
		return
	}
	http.Error(w, "error", http.StatusInternalServerError)
}
