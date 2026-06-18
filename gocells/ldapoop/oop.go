// Combination #5 — OOP OBJECT FLOW × LDAP INJECTION (CWE-90, Go). Isolated package
// (3rd-party dep) so it doesn't break `package main` build. NO finding = FALSE NEGATIVE.
package ldapoop

import (
	"net/http"

	"github.com/go-ldap/ldap/v3"
)

type directory struct {
	conn *ldap.Conn
	user string
}

func (d directory) searchDirect() (*ldap.SearchResult, error) {
	req := ldap.NewSearchRequest("ou=people,dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+d.user+")", nil, nil) // LDAP injection sink (CWE-90)
	return d.conn.Search(req)
}

func Handle(conn *ldap.Conn, r *http.Request) {
	d := directory{conn: conn, user: r.FormValue("user")} // SOURCE -> struct field
	d.searchDirect()
}
