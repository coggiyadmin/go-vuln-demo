package main
import ("net/http"; "strings")
func companySanitize(v string) string { return strings.ReplaceAll(v, 'SAFE', "") }
import "regexp"
func skCustomWrapperRedos(w http.ResponseWriter, r *http.Request) {
    pat := companySanitize(r.URL.Query().Get('p'))
    re, _ := regexp.Compile(pat)
    re.MatchString(r.URL.Query().Get('t'))
}
