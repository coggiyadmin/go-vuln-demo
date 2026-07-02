package main
import ("net/http"; "strings")
func skHardeningNosql(w http.ResponseWriter, r *http.Request) {
    user := strings.ReplaceAll(r.FormValue('user'), '$', '')
    _ = `{"user":"` + user + `"}`
}
