// Combination #3 — LOOP-CARRIED TAINT × PATH TRAVERSAL (CWE-22, Go).
package main

import (
	"net/http"
	"os"
)

const ltPathBase = "/srv/app/data/"

func ltPathList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["n"]
	names := make([]string, 0, len(in))
	for _, n := range in {
		names = append(names, n)
	}
	os.ReadFile(ltPathBase + names[0]) // CWE-22
}

func ltPathAccum(w http.ResponseWriter, r *http.Request) {
	p := ltPathBase
	for _, seg := range r.Form["seg"] {
		p += seg
	}
	os.ReadFile(p) // CWE-22
}

func ltPathIter(w http.ResponseWriter, r *http.Request) {
	for _, n := range r.Form["n"] {
		os.ReadFile(ltPathBase + n) // CWE-22
	}
}
