package closed
import (
  "database/sql"
  "net/http"
)
// cognium-dev#583 — QueryContext is a sql sink
func Lookup(db *sql.DB, r *http.Request) {
  id := r.FormValue("id")
  db.QueryContext(r.Context(), "SELECT * FROM u WHERE id="+id) // SINK CWE-89
}
