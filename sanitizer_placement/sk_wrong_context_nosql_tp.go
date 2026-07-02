package main
import ("fmt"; "html"; "net/http")
func skWrongContextNosql(w http.ResponseWriter, r *http.Request) {
    user := html.EscapeString(r.FormValue("user"))
    _ = fmt.Sprintf(`{"user":"%s"}`, user)
}
