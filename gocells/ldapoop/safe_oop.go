// SAFE mirror — ldapoop; filter metacharacters escaped via ldap.EscapeFilter. Expect 0.
package ldapoop

import (
	"net/http"

	"github.com/go-ldap/ldap/v3"
)

type safeDirectory struct {
	conn *ldap.Conn
	user string
}

func (d safeDirectory) searchDirect() (*ldap.SearchResult, error) {
	req := ldap.NewSearchRequest("ou=people,dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+ldap.EscapeFilter(d.user)+")", nil, nil) // escaped before sink
	return d.conn.Search(req)
}

func SafeHandle(conn *ldap.Conn, r *http.Request) {
	d := safeDirectory{conn: conn, user: r.FormValue("user")}
	d.searchDirect()
}
