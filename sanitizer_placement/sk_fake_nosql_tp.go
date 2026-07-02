package main
import ("fmt"; "net/http"; "strings")
func skFakeNosql(w http.ResponseWriter, r *http.Request) {
    user := strings.ReplaceAll(r.FormValue("user"), "SAFE", "")
    _ = fmt.Sprintf(`{"user":"%s"}`, user)
}
