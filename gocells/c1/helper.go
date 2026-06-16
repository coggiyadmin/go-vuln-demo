package c1
import ("database/sql";"fmt")
var cdb *sql.DB
func Run(name string){ q := fmt.Sprintf("SELECT * FROM t WHERE n='%s'", name); cdb.Query(q) } // SINK CWE-89
