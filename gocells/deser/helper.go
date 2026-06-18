package deserxf

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

func Run(blob string) (map[string]interface{}, error) {
	raw, _ := base64.StdEncoding.DecodeString(blob)
	var v map[string]interface{}
	err := gob.NewDecoder(bytes.NewReader(raw)).Decode(&v) // SINK CWE-502
	return v, err
}
