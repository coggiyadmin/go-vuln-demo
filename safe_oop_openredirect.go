// SAFE mirror — oop_openredirect.go; the tainted target is constrained to a
// relative same-site path before redirect. Expect 0 security findings.
package main

import (
	"net/http"
	"net/url"
	"strings"
)

type safeRedirNav struct{ url string }

func (n safeRedirNav) safeDest() string {
	u, err := url.Parse(n.url)
	if err != nil || u.IsAbs() || u.Host != "" || strings.HasPrefix(n.url, "//") {
		return "/" // reject absolute / cross-host targets
	}
	return "/" + strings.TrimLeft(u.Path, "/")
}

func (n safeRedirNav) goDirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, n.safeDest(), http.StatusFound)
}

func safeOpenRedirectOOPHandler(w http.ResponseWriter, r *http.Request) {
	n := safeRedirNav{url: r.FormValue("next")}
	n.goDirect(w, r)
}
