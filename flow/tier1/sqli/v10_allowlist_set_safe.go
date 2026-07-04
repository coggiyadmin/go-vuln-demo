// v10 allowlist-set — prepared stmt after int allowlist gate.
package main

import "database/sql"

var safeIDs = map[int]bool{1: true, 2: true, 3: true}

func lookup(id int) error {
	if safeIDs[id] {
		db, _ := sql.Open("sqlite3", ":memory:")
		_, err := db.Exec("SELECT * FROM users WHERE id=?", id)
		return err
	}
	return nil
}
