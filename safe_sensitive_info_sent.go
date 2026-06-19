// SAFE mirror — sensitive_info_sent.go; only non-sensitive fields are projected into the DTO.
package main

import (
	"encoding/json"
	"net/http"
)

type userDTO201 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func safeSensitiveSent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userDTO201{ID: 7, Name: "ada"}) // secrets excluded
}
