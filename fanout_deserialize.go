// Combination #11 — FAN-OUT × DESERIALIZE (CWE-502, Go).
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
)

func fanoutDeserializeHandler(w http.ResponseWriter, r *http.Request) {
	s := r.FormValue("s")
	raw, _ := base64.StdEncoding.DecodeString(s)
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&map[string]interface{}{}) // CWE-502
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&map[string]interface{}{}) // CWE-502
	gob.NewDecoder(bytes.NewReader([]byte(s))).Decode(&map[string]interface{}{}) // CWE-502
}
