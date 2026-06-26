package main

import "database/sql"

func lookup(db *sql.DB, userID int) (string, error) {
	row := db.QueryRow("SELECT name FROM users WHERE id = ?", userID)
	var name string
	return name, row.Scan(&name)
}
