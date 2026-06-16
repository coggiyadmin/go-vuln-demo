// DEMO FILE — intentional vulnerabilities for security scanner showcase (Go).
//
// CWE findings : CWE-89 (SQL Injection), CWE-78 (OS Command Injection),
//                CWE-22 (Path Traversal), CWE-918 (SSRF), CWE-79 (XSS),
//                CWE-502 (Unsafe Deserialization), CWE-798 (Hardcoded Secrets)
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

// CWE-798: hardcoded production credentials
const (
	dbPassword = "Pr0dDB@dmin!2024xyz"
	awsKey     = "AKIAIOSFODNN7EXAMPLE"
	stripeKey  = "sk_live_51ABCDEFghijklmnopqrstuvwxyz1234567890QRST"
	jwtSecret  = "hardcoded-jwt-signing-secret-never-do-this"
)

var db *sql.DB

// CWE-89: SQL Injection — user input concatenated into query string
func getUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)
	rows, err := db.Query(query) // sink: tainted SQL
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()
	w.Write([]byte("ok"))
}

// CWE-78: OS Command Injection — user input passed to shell
func ping(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("sh", "-c", "ping -c 3 "+host).Output() // sink: shell injection
	w.Write(out)
}

// CWE-22: Path Traversal — user input used directly in file path
func readFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	data, err := os.ReadFile("/var/app/uploads/" + name) // sink: traversal via ../
	if err != nil {
		http.Error(w, "not found", 404)
		return
	}
	w.Write(data)
}

// CWE-918: SSRF — user input controls outbound request target
func fetch(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	resp, err := http.Get(target) // sink: SSRF to internal services / cloud metadata
	if err != nil {
		http.Error(w, "bad gateway", 502)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	w.Write(body)
}

// CWE-79: XSS — user input written to response without encoding
func greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "<h1>Hello %s</h1>", name) // sink: reflected XSS
}

// CWE-502: Unsafe Deserialization — untrusted JSON into an interface{} target
func importConfig(w http.ResponseWriter, r *http.Request) {
	raw, _ := io.ReadAll(r.Body)
	var cfg map[string]interface{}
	json.Unmarshal(raw, &cfg) // sink: deserialization of untrusted data
	fmt.Fprintf(w, "%v", cfg)
}

func main() {
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/file", readFile)
	http.HandleFunc("/fetch", fetch)
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/import", importConfig)
	http.ListenAndServe(":8080", nil)
}
