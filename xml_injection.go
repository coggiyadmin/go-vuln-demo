// CWE-91 — XML Injection (Go). Unescaped input concatenated into an XML document injects
// structure. NO finding = FALSE NEGATIVE.
package main

import "net/http"

func xmlInjOrder(w http.ResponseWriter, r *http.Request) {
	qty := r.FormValue("qty") // SOURCE
	xml := "<order><qty>" + qty + "</qty><price>100</price></order>" // injects elements → CWE-91
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml))
}
