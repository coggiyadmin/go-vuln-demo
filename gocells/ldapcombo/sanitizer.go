// Combinations #6/#7/#8 — SANITIZER × LDAP (CWE-90). Isolated package.
package ldapcombo

import (
	"net/http"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func escapeHtml(s string) string { return strings.ReplaceAll(s, "<", "_") }
func sanitize(s string) string   { return s }
func ldapSafe(s string) string {
	return strings.NewReplacer("(", "", ")", "", "*", "", "\\", "").Replace(s)
}

func Wrong(conn *ldap.Conn, r *http.Request) {
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+escapeHtml(r.FormValue("user"))+")", nil, nil)) // CWE-90
}

func Fake(conn *ldap.Conn, r *http.Request) {
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+sanitize(r.FormValue("user"))+")", nil, nil)) // CWE-90
}

func Wrapped(conn *ldap.Conn, r *http.Request) {
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+ldapSafe(r.FormValue("user"))+")", nil, nil)) // expect 0
}
