package profile

import "database/sql"

// TN — library API with bound query (#169).
func Lookup(db *sql.DB, id int) (string, error) {
    row := db.QueryRow("SELECT name FROM users WHERE id = ?", id)
    var name string
    return name, row.Scan(&name)
}
