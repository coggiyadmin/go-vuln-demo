// FN probe — CWE-201 Insertion of Sensitive Information Into Sent Data. The full user record,
// including the password hash and API token, is serialized to the client. NO finding = FN.
package main

import (
	"encoding/json"
	"net/http"
)

type userRecord201 struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"passwordHash"`
	APIToken     string `json:"apiToken"`
}

func sensitiveSent(w http.ResponseWriter, r *http.Request) {
	u := userRecord201{ID: 7, Name: "ada", PasswordHash: "$2b$12$abcdef", APIToken: "sk-live-9931"}
	json.NewEncoder(w).Encode(u) // leaks passwordHash + apiToken → CWE-201
}
