package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, 'pickle', "") }
func skCustomWrapperDeserialize(w http.ResponseWriter, r *http.Request) {
    _ = companySanitize(r.FormValue('blob'))
}
