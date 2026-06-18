package ldapxf

import (
	"net/http"

	"github.com/go-ldap/ldap/v3"
)

func Handle(conn *ldap.Conn, r *http.Request) { Run(conn, r.FormValue("user")) } // SOURCE -> cross-file sink (CWE-90)
