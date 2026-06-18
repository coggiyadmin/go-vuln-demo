package ldapxf

import "github.com/go-ldap/ldap/v3"

func Run(conn *ldap.Conn, user string) (*ldap.SearchResult, error) {
	req := ldap.NewSearchRequest("ou=people,dc=example,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+user+")", nil, nil) // SINK CWE-90
	return conn.Search(req)
}
