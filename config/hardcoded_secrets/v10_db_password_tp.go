package main
import ("database/sql")
func connect() (*sql.DB, error) {
    return sql.Open("postgres", "host=db user=app password=S3cr3tP@ss! dbname=app") // SINK CWE-798
}
