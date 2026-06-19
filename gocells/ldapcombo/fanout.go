package ldapcombo

import (
	"net/http"

	"github.com/go-ldap/ldap/v3"
)

func Fanout(conn *ldap.Conn, r *http.Request) {
	u := r.FormValue("u")
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+u+")", nil, nil)) // CWE-90
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(cn="+u+")", nil, nil)) // CWE-90
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(mail="+u+")", nil, nil)) // CWE-90
}
