// FP-target (upstream cognium-dev#128/#140) — DEEP real-*sql.DB* flow. A repository helper
// builds SQL from a caller-supplied `filter` and runs it on a real *sql.DB sink (QueryContext).
// `filter` is library-API input from the integrating app, NOT an HTTP entry point, so under an
// entry-point gate this must NOT be sql_injection — even though a genuine concat→Query sink is
// present. (Confirms the engine isn't firing purely on the sink shape with no source.)
package profile

import (
	"context"
	"database/sql"
)

// ItemRepo is a library data-access helper embedded by applications.
type ItemRepo struct{ db *sql.DB }

// FindWhere runs a caller-supplied filter clause. `filter` originates from the embedding
// application's own config/DSL, not from request data.
func (r *ItemRepo) FindWhere(ctx context.Context, filter string) (*sql.Rows, error) {
	query := "SELECT id, name FROM items WHERE " + filter // real sink, library-supplied operand
	return r.db.QueryContext(ctx, query)
}
