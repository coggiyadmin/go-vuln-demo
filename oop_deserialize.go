// Combination #5 — OOP OBJECT FLOW × INSECURE DESERIALIZATION (CWE-502, Go). Tainted
// struct field gob-decoded (untrusted). NO finding = FALSE NEGATIVE.
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
)

type session struct{ blob string }

func (s session) loadDirect() (map[string]interface{}, error) {
	raw, _ := base64.StdEncoding.DecodeString(s.blob)
	var v map[string]interface{}
	err := gob.NewDecoder(bytes.NewReader(raw)).Decode(&v) // deserialize sink (CWE-502)
	return v, err
}

func deserializeOOPHandler(w http.ResponseWriter, r *http.Request) {
	s := session{blob: r.FormValue("s")} // SOURCE -> struct field
	s.loadDirect()
}
