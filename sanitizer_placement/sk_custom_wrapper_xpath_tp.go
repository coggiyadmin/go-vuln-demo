package main
import ("net/http"; "strings")
func companySanitize(v string) string {
    return strings.NewReplacer("'", "", `"`, "").Replace(v)
}
func skCustomWrapperXpath(w http.ResponseWriter, r *http.Request) {
    name := companySanitize(r.URL.Query().Get("name"))
    _ = "//user[name='" + name + "']"
}
