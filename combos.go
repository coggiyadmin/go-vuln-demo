// Combination × language matrix — GO row (self-contained: own *sql.DB + var
// pattern, so neither cross-file #74 nor inline-Sprintf confounds the result).
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var cdb *sql.DB // local sink receiver (typed *sql.DB)

// #2 path-sensitivity — tainted used in the failure branch
func c2Path(w http.ResponseWriter, r *http.Request) {
	x := r.FormValue("x")
	if len(x) > 100 {
		q := fmt.Sprintf("SELECT * FROM t WHERE n='%s'", x)
		cdb.Query(q) // CWE-89
	}
}

// #3 loop-carried — iterate over a tainted multi-value param
func c3Loop(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()["p"]
	for _, a := range vals {
		q := fmt.Sprintf("SELECT * FROM t WHERE n='%s'", a)
		cdb.Query(q) // CWE-89
	}
}

// #5 struct/OOP flow — field set from input, used in a method
type Job struct{ q string }

func (j Job) run() {
	s := fmt.Sprintf("SELECT %s", j.q)
	cdb.Query(s) // CWE-89
}
func c5OOP(w http.ResponseWriter, r *http.Request) {
	j := Job{q: r.FormValue("q")}
	j.run()
}

// #11 fan-out — one source → two sinks
func c11Fanout(w http.ResponseWriter, r *http.Request) {
	u := r.FormValue("u")
	q := fmt.Sprintf("SELECT %s", u)
	cdb.Query(q) // sink 1 (sql)
	p := fmt.Sprintf("/var/app/%s", u)
	_, _ = os.ReadFile(p) // sink 2 (path)
}

// #7 fake sanitizer — no-op
func clean(x string) string { return x }
func c7Fake(w http.ResponseWriter, r *http.Request) {
	u := clean(r.FormValue("u"))
	q := fmt.Sprintf("SELECT %s", u)
	cdb.Query(q) // CWE-89
}
