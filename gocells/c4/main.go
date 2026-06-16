package main
import ("database/sql";"fmt";"net/http")
var cdb *sql.DB
func C4(r *http.Request){ c := r.FormValue("c"); q := fmt.Sprintf("SELECT %s", c); go func(){ cdb.Query(q) }() } // CWE-89 via goroutine
func main(){}
