// FN probe — Go string concatenation drops taint (#53).
// "literal" + taintedVar must propagate to exec.Command and os.ReadFile sinks.
package main

import (
	"net/http"
	"os"
	"os/exec"
)

func pingConcat(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("sh", "-c", "ping -c 3 "+host).Output() // CWE-78 @14
	w.Write(out)
}

func readConcat(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	data, err := os.ReadFile("/var/app/uploads/" + name) // CWE-22 @20
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.Write(data)
}
