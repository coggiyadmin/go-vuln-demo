package ldapcombo

import (
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/go-ldap/ldap/v3"
)

func EncodedB64(conn *ldap.Conn, r *http.Request) {
	u, _ := base64.StdEncoding.DecodeString(r.FormValue("d"))
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+string(u)+")", nil, nil)) // CWE-90
}

func EncodedURL(conn *ldap.Conn, r *http.Request) {
	u, _ := url.QueryUnescape(r.FormValue("d"))
	conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(uid="+u+")", nil, nil)) // CWE-90
}
