// FN probe — CWE-477 Use of Obsolete Function. `ioutil.ReadAll` is deprecated since Go 1.16
// (use io.ReadAll). Real vuln; NO finding = FN.
package main

import (
	"io/ioutil"
	"net/http"
)

func obsoleteRead(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body) // obsolete/deprecated API → CWE-477
	w.Write(body)
}
