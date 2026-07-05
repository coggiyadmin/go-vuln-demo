package attacksurface
import ("database/sql"; "os"; "testing")
func TestQuery(t *testing.T) {
    q := os.Getenv("Q")
    var db *sql.DB
    db.Query("SELECT * FROM u WHERE n='" + q + "'") // SINK — test file
}
