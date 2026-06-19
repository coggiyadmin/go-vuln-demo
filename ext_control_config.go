// FN probe — CWE-15 External Control of System or Configuration Setting.
// A user-controlled key/value writes directly into the process environment.
package main

import (
	"net/http"
	"os"
)

func extControlConfig(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("k") // SOURCE — attacker-controlled setting name
	val := r.URL.Query().Get("v")
	os.Setenv(key, val) // user controls arbitrary config → CWE-15
}
