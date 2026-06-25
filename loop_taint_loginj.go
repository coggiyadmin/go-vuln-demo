// Combination #3 — LOOP-CARRIED TAINT × LOGINJ (CWE-117, Go). Taint flows through a
// loop into the sink. A handler with NO finding is a FALSE NEGATIVE.
package main

import (
	"log"
	"net/http"
)

// 3a. SLICE ELEMENT BUILT IN LOOP
func ltLoList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["user"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	user := items[0]
	log.Printf("login user=%s", user) // CWE-117
}

// 3b. STRING ACCUMULATOR
func ltLoAccum(w http.ResponseWriter, r *http.Request) {
	user := ""
	for _, x := range r.Form["user"] { user += x }
	log.Printf("login user=%s", user) // CWE-117
}

// 3c. ITERATE-AND-SINK
func ltLoIter(w http.ResponseWriter, r *http.Request) {
	for _, user := range r.Form["user"] {
		log.Printf("login user=%s", user) // CWE-117 per iteration
	}
}
