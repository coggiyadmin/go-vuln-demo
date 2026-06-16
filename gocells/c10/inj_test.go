package main
import ("database/sql";"fmt";"net/http")
var cdb *sql.DB
func Inj(r *http.Request){ c := r.FormValue("c"); q := fmt.Sprintf("SELECT %s", c); cdb.Query(q) } // CWE-89
