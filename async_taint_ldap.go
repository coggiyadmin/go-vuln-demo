// Combination #4 — ASYNC taint × LDAP (CWE-90, Go).
package main

import (
	"net/http"
	"github.com/go-ldap/ldap/v3"
)

func atLdapHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	ch := make(chan string, 1)
	go func() { ch <- uid }()
	v := <-ch
	conn, _ := ldap.DialURL("ldap://localhost")
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+uid+")", nil, nil)) // CWE-90
}
