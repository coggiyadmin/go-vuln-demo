package main
import ("net/http"; "github.com/go-ldap/ldap/v3")
func ldapFilter(w http.ResponseWriter, r *http.Request) {
    uid := r.FormValue("uid")
    conn, _ := ldap.DialURL("ldap://localhost")
    conn.Search(ldap.NewSearchRequest("ou=people", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
        "(uid="+uid+")", nil, nil)) // SINK CWE-90
}
