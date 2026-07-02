// SAFE — constant query only.
package interop

func LookupEnvJsonSqlSafe() {
    _ = "SELECT * FROM u WHERE n='alice'"
}
