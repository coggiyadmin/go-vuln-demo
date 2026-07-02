package main
import ("database/sql"; "fmt"; "net/http"; _ "github.com/mattn/go-sqlite3")
func skFrameworkNativeSqli(w http.ResponseWriter, r *http.Request) {
    n := r.URL.Query().Get("n")
    db, _ := sql.Open("sqlite3", ":memory:")
    db.Query(fmt.Sprintf("SELECT * FROM u WHERE n='%s'", n))
}
