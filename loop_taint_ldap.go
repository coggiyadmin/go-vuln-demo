// Combination #3 — LOOP-CARRIED TAINT × LDAP (CWE-90, Go).
package main

import (
	"net/http"
	"github.com/go-ldap/ldap/v3"
)

func ltLdapList(w http.ResponseWriter, r *http.Request) {
	in := r.Form["uid"]
	items := make([]string, 0, len(in))
	for _, x := range in { items = append(items, x) }
	uid := items[0]
	conn, _ := ldap.DialURL("ldap://localhost")
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+uid+")", nil, nil)) // CWE-90
}
