// Combinations #6/#7/#8 — SANITIZER × DESERIALIZE (CWE-502, Go).
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"strings"
)

func scStripHTMLDeser(s string) string { return strings.NewReplacer("<", "_").Replace(s) }
func scFakeDeser(s string) string       { return s }

func scWrongDeserialize(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(scStripHTMLDeser(r.FormValue("s")))
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&map[string]interface{}{}) // CWE-502
}

func scFakeDeserialize(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(scFakeDeser(r.FormValue("s")))
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&map[string]interface{}{}) // CWE-502
}

func scWrappedDeserialize(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.FormValue("s"))
	json.Unmarshal(raw, &map[string]interface{}{}) // expect 0
}
