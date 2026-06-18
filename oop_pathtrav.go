// Combination #5 — OOP OBJECT FLOW × PATH TRAVERSAL (CWE-22, Go). Taint stored in
// a struct field, read in a method sink. NO finding = FALSE NEGATIVE.
package main

import (
	"net/http"
	"os"
)

const pathBase = "/srv/app/data/"

type pathDoc struct{ name string }

func (d pathDoc) readDirect() ([]byte, error) {
	return os.ReadFile(pathBase + d.name) // path-traversal sink (CWE-22)
}

func pathTravOOPHandler(w http.ResponseWriter, r *http.Request) {
	d := pathDoc{name: r.FormValue("name")} // SOURCE -> struct field
	d.readDirect()
}
