package main
import ("database/sql")
// password supplied via an already-resolved DSN from the secrets layer (no literal)
func connect(dsn string) (*sql.DB, error) {
    return sql.Open("postgres", dsn)
}
