package main
import ("fmt"; "net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, "$", "") }
func skCustomWrapperNosql(w http.ResponseWriter, r *http.Request) {
    user := companySanitize(r.FormValue("user"))
    _ = fmt.Sprintf(`{"user":"%s"}`, user)
}
