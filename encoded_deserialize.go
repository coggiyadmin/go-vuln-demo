// Combination #13 — ENCODED × DESERIALIZE (CWE-502, Go).
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
	"net/url"
)

func encodedDeserializeB64Handler(w http.ResponseWriter, r *http.Request) {
	raw, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&map[string]interface{}{}) // CWE-502
}

func encodedDeserializeURLHandler(w http.ResponseWriter, r *http.Request) {
	s, _ := url.QueryUnescape(r.FormValue("d"))
	gob.NewDecoder(bytes.NewReader([]byte(s))).Decode(&map[string]interface{}{}) // CWE-502
}
