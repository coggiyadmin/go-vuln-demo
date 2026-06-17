// NEGATIVE TEST FILE — safe mirror of init_harvest.go.
//
// The scanner MUST produce ZERO security findings here. Any finding is a
// FALSE POSITIVE.
package main

import (
	"log"
	"os"
	"time"
)

func init() {
	// SAFE init — no credential reads, no network exfiltration.
	go logStartup()
}

func logStartup() {
	log.Printf("service started pid=%d", os.Getpid())
	time.Sleep(100 * time.Millisecond)
}
