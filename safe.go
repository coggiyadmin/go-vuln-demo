// NEGATIVE TEST FILE — secure equivalents of every vulnerable pattern (Go).
//
// Flows user input through safe APIs to each sink type. The scanner MUST
// produce ZERO security findings here. Any finding is a FALSE POSITIVE.
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const uploadRoot = "/var/app/uploads"

var allowedHosts = map[string]bool{
	"api.internal.example.com": true,
	"cdn.example.com":          true,
}

// SAFE sql — parameterized query with placeholder, no Sprintf
func safeGetUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username) // parameterized
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}

// SAFE command — fixed program + arg slice, no shell
func safePing(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("ping", "-c", "3", "--", host).Output() // no shell, host is one argv
	w.Write(out)
}

// SAFE path — clean + prefix containment check
func safeReadFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	clean := filepath.Clean(filepath.Join(uploadRoot, name))
	if !strings.HasPrefix(clean, uploadRoot+string(os.PathSeparator)) {
		http.Error(w, "forbidden", 403)
		return
	}
	data, err := os.ReadFile(clean)
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.Write(data)
}

// SAFE ssrf — host validated against allowlist before request
func safeFetch(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	parsed, err := parseAndCheck(target)
	if err != nil {
		http.Error(w, "host not allowed", 403)
		return
	}
	resp, err := http.Get(parsed)
	if err != nil {
		http.Error(w, "bad gateway", 502)
		return
	}
	defer resp.Body.Close()
	w.Write([]byte("ok"))
}

func parseAndCheck(raw string) (string, error) {
	u, err := neturlParse(raw)
	if err != nil {
		return "", err
	}
	if !allowedHosts[u] {
		return "", fmt.Errorf("host not allowed")
	}
	return raw, nil
}

// SAFE config — credentials from environment, nothing hardcoded
func dbURL() string {
	return os.Getenv("DATABASE_URL")
}

// helper kept simple for the demo
func neturlParse(raw string) (string, error) {
	parts := strings.SplitN(strings.TrimPrefix(strings.TrimPrefix(raw, "https://"), "http://"), "/", 2)
	if len(parts) == 0 {
		return "", fmt.Errorf("bad url")
	}
	return parts[0], nil
}
