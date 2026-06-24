// FP-target (cognium-dev#163, go) — SQL *identifier* (column) validated against an allowlist
// regex and quoted; the value is bound as a placeholder parameter. Must not be flagged
// sql_injection.
package fpcorpus

import (
	"database/sql"
	"fmt"
	"regexp"
)

var identRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func quoteIdent(id string) (string, error) {
	if !identRe.MatchString(id) {
		return "", fmt.Errorf("bad identifier")
	}
	return `"` + id + `"`, nil
}

// ByColumn validates+quotes the column identifier and binds the value as a parameter.
func ByColumn(db *sql.DB, column, value string) (*sql.Rows, error) {
	col, err := quoteIdent(column)
	if err != nil {
		return nil, err
	}
	return db.Query("SELECT * FROM items WHERE "+col+" = ?", value) // ident validated, value bound
}
