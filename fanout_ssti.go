// Combination #11 — FAN-OUT × SSTI (CWE-1336, Go).
package main

import (
	"html/template"
	"net/http"
)

func fanoutSstiHandler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("n")
	t1 := template.Must(template.New("t1").Parse("<p>" + n + "</p>"))
	t1.Execute(w, nil) // CWE-1336
	t2 := template.Must(template.New("t2").Parse("<h1>" + n + "</h1>"))
	t2.Execute(w, nil) // CWE-1336
	t3 := template.Must(template.New("t3").Parse("{{.N}}"))
	t3.Execute(w, map[string]string{"N": n}) // CWE-1336
}
