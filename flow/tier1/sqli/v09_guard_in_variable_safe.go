// v09 guard-in-variable — column allowlist in package map.
package main

import ("database/sql"; "fmt")

var allowedCols = map[string]bool{"name": true, "email": true}

func lookup(col, val string) error {
	if allowedCols[col] {
		db, _ := sql.Open("sqlite3", ":memory:")
		_, err := db.Exec(fmt.Sprintf("SELECT %s FROM users WHERE %s=?", col, col), val)
		return err
	}
	return nil
}
