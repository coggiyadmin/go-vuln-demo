// SAFE mirror — sensitive_info_debug.go; the debug branch returns only a correlation id.
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func safeSensitiveDebug(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("boom")
	b := make([]byte, 8)
	rand.Read(b)
	ref := hex.EncodeToString(b)
	log.Printf("op failed ref=%s err=%v", ref, err) // detail stays server-side
	fmt.Fprintf(w, "error ref=%s", ref)
}
