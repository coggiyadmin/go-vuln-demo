// SAFE mirror — xml_injection.go; value XML-escaped before insertion. Expect 0 findings.
package main

import (
	"net/http"
	"strings"
)

func safeXmlInjOrder(w http.ResponseWriter, r *http.Request) {
	qty := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(r.FormValue("qty")) // escaped
	xml := "<order><qty>" + qty + "</qty><price>100</price></order>"
	w.Write([]byte(xml))
}
