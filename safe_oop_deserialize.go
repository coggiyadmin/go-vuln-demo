// SAFE mirror — oop_deserialize.go; typed JSON into a fixed struct, no arbitrary
// type instantiation. Expect 0 security findings.
package main

import (
	"encoding/json"
	"net/http"
)

type safeSession struct{ blob string }

type sessData struct {
	User string `json:"user"`
}

func (s safeSession) loadDirect() (sessData, error) {
	var d sessData
	err := json.Unmarshal([]byte(s.blob), &d) // typed JSON, no code execution
	return d, err
}

func safeDeserializeOOPHandler(w http.ResponseWriter, r *http.Request) {
	s := safeSession{blob: r.FormValue("s")}
	s.loadDirect()
}
