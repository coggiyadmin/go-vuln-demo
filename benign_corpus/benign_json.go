package main

import "encoding/json"

func parse(raw []byte) (map[string]any, error) {
	var out map[string]any
	return out, json.Unmarshal(raw, &out)
}
