// SAFE mirror — obsolete_function.go; uses the modern `io.ReadAll` replacement.
package main

import (
	"io"
	"net/http"
)

func safeObsoleteRead(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(io.LimitReader(r.Body, 1<<20)) // modern, bounded
	w.Write(body)
}
