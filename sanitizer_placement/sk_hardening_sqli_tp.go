package main
import ("net/http"; "strings")
func skHardeningSqli(w http.ResponseWriter, r *http.Request) {
    n := strings.ReplaceAll(r.URL.Query().Get('n'), '--', '')
    _ = "SELECT * FROM u WHERE n='" + n + "'"
}
