// FN probe — CWE-538 Insertion of Sensitive Information into Externally-Accessible File.
// A secret is written to a web-served static directory.
package main

import (
	"net/http"
	"os"
)

func exportConfig(w http.ResponseWriter, r *http.Request) {
	// writes the API key into a publicly-served static path → CWE-538
	os.WriteFile("/var/www/static/config.txt", []byte("API_KEY="+os.Getenv("API_KEY")), 0o644)
}
